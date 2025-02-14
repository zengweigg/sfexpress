package sfexpress

import (
	"fmt"
	"github.com/zengweigg/sfexpress/model"
)

type icmsService service

// GetCityList 查询城市列表
func (s icmsService) GetCityList(postData model.SFICMSQueryCityList) (model.SFICMSQueryCityResp, error) {
	respData := model.SFICMSQueryCityResp{}
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_CITY_LIST")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}

	// 解码数据
	//res := model.SFResponse{}
	//err = jsoniter.Unmarshal(resp.Body(), &res)
	//if err != nil {
	//	return respData, nil
	//}
	//ok, err := DecodeMsg(res.ApiResultData, token, aeskey, apikey)
	//fmt.Println("ok", ok)
	//// 把字符串转换为结构体
	//if err == nil {
	//	err = jsoniter.Unmarshal([]byte(ok), &respData)
	//}
	//fmt.Println("respData", respData)
	return respData, nil
}
