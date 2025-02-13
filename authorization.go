package sfexpress

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	"net/url"
	"sfexpress/config"
	"time"
)

type authorizationService service

func newHttpClient(c config.Config) *resty.Client {
	httpClient := resty.
		New().
		SetDebug(c.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
			"lang":         c.Lang,
		})
	if c.Sandbox {
		httpClient.SetBaseURL("http://api-ifsp-sit.sf.global")
	} else {
		httpClient.SetBaseURL("https://api-ifsp.sf.global")
	}
	return httpClient
}

// GetToken 获取 access-token
// https://openapi-portal.sf.global/document?firstMenu=DOC&secondMenu=75
func (s authorizationService) GetToken() (ar Token, err error) {
	// 定义一个结构体用于解析API返回的JSON数据
	result := struct {
		Code         int    `json:"apiResultCode"` // API返回的结果代码
		Message      string `json:"apiErrorMsg"`   // API返回的错误信息
		Data         Token  `json:"apiResultData"` // API返回的Token数据
		ApiTimestamp int64  `json:"apiTimestamp"`  // API返回的时间戳
	}{}

	// 创建一个新的HTTP客户端，并设置日志记录器
	// 设置返回结果解析到result结构体
	resp, err := newHttpClient(*s.config).
		SetLogger(s.logger).
		R().
		SetResult(&result).
		Get(fmt.Sprintf("/openapi/api/token?appKey=%s&appSecret=%s", s.config.APIKey, url.QueryEscape(s.config.APISecret))) // 发送GET请求获取Token

	// 如果请求过程中发生错误，返回错误信息
	if err != nil {
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
		return
	}

	// 如果请求成功
	if resp.IsSuccess() {
		code := result.Code
		// 如果API返回的结果代码不为0，表示有错误，返回错误信息
		if code != 0 {
			err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
			return
		}
		// 将API返回的Token数据赋值给ar
		ar = result.Data
		// 计算Token的过期时间，设置为当前时间加上Token有效期的90%（剩余10%时间需要更换Token）
		ar.ExpiresDatetime = time.Now().Unix() + int64(ar.ExpiresIn*9/10)
	} else {
		// 如果请求不成功，返回错误信息
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
	}
	return
}
