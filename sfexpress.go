package sfexpress

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/config"
	"github.com/zengweigg/sfexpress/model"
	"net/http"
	"strconv"
	"time"
)

const (
	Version   = "1.0.0"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"
)

type SFClient struct {
	config     *config.Config
	httpClient *resty.Client
	logger     Logger   // Logger
	Services   services // API Services
}

func NewSFService(cfg config.Config) *SFClient {
	SFClient := &SFClient{
		config: &cfg,
		logger: createLogger(),
	}
	//OnBeforeRequest：设置请求发送前的钩子函数，允许在请求发送之前对请求进行修改或添加逻辑。
	//OnAfterResponse：设置响应接收后的钩子函数，允许在接收到响应后处理响应数据或执行其他逻辑。
	//SetRetryCount：设置请求失败时的最大重试次数。
	//SetRetryWaitTime：设置每次重试之间的等待时间（最小等待时间）。
	//SetRetryMaxWaitTime：设置每次重试之间的最大等待时间，实际等待时间会在最小和最大等待时间之间随机选取。
	//AddRetryCondition：添加自定义的重试条件，当满足该条件时触发重试机制。
	httpClient := resty.
		New().
		SetDebug(SFClient.config.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	if cfg.Sandbox {
		httpClient.SetBaseURL("http://api-ifsp-sit.sf.global/openapi/api/dispatch")
	} else {
		httpClient.SetBaseURL("https://api-ifsp.sf.global/openapi/api/dispatch")
	}
	var tempToken string
	httpClient.
		SetTimeout(time.Duration(cfg.Timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// 获取token
			ft := FileToken{}
			var tokenWriterReader TokenWriterReader = ft
			tokens, err := tokenWriterReader.GetValidAccessToken(cfg.APIKey, cfg.APISecret, cfg.Lang, cfg.Debug, cfg.Sandbox)
			if err != nil {
				return err
			}
			tempToken = tokens.AccessToken
			// 获取签名
			param := "{}"
			if request.Body != nil {
				bd, e := jsoniter.Marshal(request.Body)
				if e != nil {
					return e
				}
				param = string(bd)
			}
			timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
			p, err := GetSign(param, tempToken, cfg.EncodingAesKey, cfg.APIKey, timestamp)
			if err != nil {
				return err
			}
			request.SetHeaders(map[string]string{
				"msgType":   request.URL,
				"appKey":    cfg.APIKey,
				"token":     tempToken,
				"timestamp": timestamp,
				"nonce":     timestamp,
				"signature": p.Signature,
				"lang":      "zh-CN",
			})
			request.SetBody(p.Encrypt)
			request.URL = ""
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsError() {
				return fmt.Errorf("%s: %s", response.Status(), bytex.ToString(response.Body()))
			}
			contentType := response.RawResponse.Header.Get("Content-Type")
			if contentType == "application/octet-stream" {
				return
			}
			r := model.SFResponse{}
			if err = jsoniter.Unmarshal(response.Body(), &r); err == nil {
				if r.ApiResultCode != 0 {
					return fmt.Errorf("%d: %s", r.ApiResultCode, r.ApiErrorMsg)
				}
				//  自定义响应数据
				if r.ApiResultData != "" {
					ok, err := DecodeMsg(r.ApiResultData, tempToken, cfg.EncodingAesKey, cfg.APIKey)
					if err == nil {
						response.SetBody([]byte(ok))
					}
				}
			} else {
				SFClient.logger.Errorf("JSON Unmarshal error: %s", err.Error())
			}
			if err != nil {
				SFClient.logger.Errorf("OnAfterResponse error: %s", err.Error())
			}
			return
		}).
		SetRetryCount(2).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			text := r.Request.URL
			if r == nil {
				return false
			}
			if err != nil {
				text += fmt.Sprintf(", error: %s", err.Error())
				SFClient.logger.Debugf("Retry request: %s", text)
				return true // 如果有错误则重试
			}
			// 检查响应状态码是否不是200
			if r.StatusCode() != http.StatusOK {
				text += fmt.Sprintf(", error: %d", r.StatusCode())
				SFClient.logger.Debugf("Retry request: %s", text)
				return true
			}
			type ResponseBody struct {
				Code int `json:"apiResultCode"`
			}
			// 解析响应体JSON
			var responseBody ResponseBody
			if err := json.Unmarshal(r.Body(), &responseBody); err != nil {
				text += fmt.Sprintf(", error: %s", string(r.Body()))
				SFClient.logger.Debugf("Retry request: %s", text)
				return true // 如果解析错误则重试
			}

			// 检查apiResultCode字段是否不是0
			if responseBody.Code != 0 {
				text += fmt.Sprintf(", error: %d", responseBody.Code)
				SFClient.logger.Debugf("Retry request: %s", text)
				return true
			}
			return false
		})
	SFClient.httpClient = httpClient
	xService := service{
		config:     &cfg,
		logger:     SFClient.logger,
		httpClient: SFClient.httpClient,
	}
	SFClient.Services = services{
		Iecs: (iecsService)(xService), // 出口电商业务场景
		//Iuop:          (iuopService)(xService), //  国际快递业务场景接入
		//Ades:          (adesservice)(xService),
		Gts:  (gtsService)(xService),  // 通用查询
		Icms: (icmsService)(xService), // 国际区划
		Icsm: (icsmService)(xService), // 关务发票
	}
	return SFClient
}
