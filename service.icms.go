package sfexpress

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/model"
)

type icmsService service

// GetContinentList 查询大洲列表
func (s icmsService) GetContinentList(postData model.SFBlank) (model.ICMSQueryContinentResp, error) {
	respData := model.ICMSQueryContinentResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("GET_CONTINENT_LIST")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// GetCityList 查询城市列表
func (s icmsService) GetCityList(postData model.ICMSQueryCityPost) (model.ICMSQueryCityResp, error) {
	respData := model.ICMSQueryCityResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_CITY_LIST")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// GetCountryList 查询国家列表
func (s icmsService) GetCountryList(postData model.ICMSCountryPost) (model.ICMSQueryCityResp, error) {
	respData := model.ICMSQueryCityResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_COUNTRY_LIST")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryCodeAttribute 地址区划-查询邮编属性
func (s icmsService) QueryCodeAttribute(postData model.ICMSQueryCodeAttributePost) (model.ICMSQueryCodeAttributeResp, error) {
	respData := model.ICMSQueryCodeAttributeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_POST_CODE_ATTRIBUTE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryCodeByRegion 地址区划-根据国际区划查询邮编
func (s icmsService) QueryCodeByRegion(postData model.ICMSQueryCodeByRegionPost) (model.ICMSQueryCodeByRegionResp, error) {
	respData := model.ICMSQueryCodeByRegionResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_POST_CODE_BY_REGION")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryCodeInfo 地址区划-查询邮编详细信息
func (s icmsService) QueryCodeInfo(postData model.ICMSQueryCodeInfoPost) (model.ICMSQueryCodeInfoResp, error) {
	respData := model.ICMSQueryCodeInfoResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_POST_CODE_INFO")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryCodeVague 地址区划-邮编模糊查询
func (s icmsService) QueryCodeVague(postData model.ICMSQueryCodeVaguePost) (model.ICMSQueryCodeVagueResp, error) {
	respData := model.ICMSQueryCodeVagueResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_POST_CODE_VAGUE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryProvinceList 地址区划-查询州省列表
func (s icmsService) QueryProvinceList(postData model.ICMSQueryProvinceListPost) (model.ICMSQueryProvinceListResp, error) {
	respData := model.ICMSQueryProvinceListResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_PROVINCE_LIST")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// QueryRegionByCode 地址区划-根据邮编查询国际区划
func (s icmsService) QueryRegionByCode(postData model.ICMSQueryRegionByCodePost) (model.ICMSQueryRegionByCodeResp, error) {
	respData := model.ICMSQueryRegionByCodeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_QUERY_REGION_BY_POST_CODE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// RegionCascadeQuery 地址区划-国际区划级联查询
func (s icmsService) RegionCascadeQuery(postData model.ICMSRegionCascadeQueryPost) (model.ICMSRegionCascadeQueryResp, error) {
	respData := model.ICMSRegionCascadeQueryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_REGION_CASCADE_QUERY")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// RegionVagueQuery 地址区划-国际区划模糊查询
func (s icmsService) RegionVagueQuery(postData model.ICMSRegionVagueQueryPost) (model.ICMSRegionVagueQueryResp, error) {
	respData := model.ICMSRegionVagueQueryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_REGION_VAGUE_QUERY")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	fmt.Println("respData", respData)
	return respData, nil
}

// PhoneValidCheck 地址区划-国际电话有效性校验
func (s icmsService) PhoneValidCheck(postData model.ICMSPhoneValidCheckPost) (model.ICMSPhoneValidCheckResp, error) {
	respData := model.ICMSPhoneValidCheckResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICMS_PHONE_VALID_CHECK")
	fmt.Println("body:", string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	fmt.Println("respData12", respData)
	return respData, nil
}
