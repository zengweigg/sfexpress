package sfexpress

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"sfexpress/model"
	"strconv"
	"strings"
	"time"
)

type icmsService service

// GetCityList 查询城市列表
func (s icmsService) GetCityList(postData model.SFICMSQueryCityList) (model.SFICMSQueryCityResp, error) {
	respData := model.SFICMSQueryCityResp{}

	// 获取token
	aeskey := s.config.EncodingAesKey
	apikey := s.config.APIKey
	var tokenWriterReader TokenWriterReader
	var auth authorizationService
	tokens, err := tokenWriterReader.GetValidAccessToken(apikey, auth)
	if err != nil {
		return respData, err
	}
	token := tokens.AccessToken

	// 转换请求数据
	ostr, _ := json.Marshal(postData)
	param := string(ostr)

	// 获取签名
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	p, err := GetSign(param, token, aeskey, apikey, timestamp)
	if err != nil {
		fmt.Println("err GetSign:", err)
		return respData, err
	}

	// 请求数据
	payload := strings.NewReader(p.Encrypt)
	resp, err := s.httpClient.R().
		SetHeaders(map[string]string{
			"msgType":   "ICMS_QUERY_CITY_LIST",
			"appKey":    apikey,
			"token":     tokens.AccessToken,
			"timestamp": timestamp,
			"nonce":     timestamp,
			"signature": p.Signature,
		}).
		SetBody(payload).
		Post("/openapi/api/dispatch")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}

	// 解码数据
	res := model.SFResponse{}
	err = jsoniter.Unmarshal(resp.Body(), &res)
	if err != nil {
		return respData, nil
	}
	ok, err := DecodeMsg(res.ApiResultData, token, aeskey, apikey)

	// 把字符串转换为结构体
	if err == nil {
		err = jsoniter.Unmarshal([]byte(ok), &respData)
	}
	return respData, nil
}
