package model

// 地址区划

// SFBlank 空白结构体
type SFBlank struct{}

// ICMSContinentData 1.查询大洲列表
type ICMSContinentData struct {
	ContinentCode     string `json:"continentCode"`     // 大洲代码
	ContinentName     string `json:"continentName"`     // 大洲名称
	ContinentNameI18n string `json:"continentNameI18n"` // 大洲名称国际化
}

type ICMSQueryContinentResp struct {
	SFApiResponse
	Data []ICMSContinentData `json:"data"`
}

// ICMSQueryCityPost 2.查询城市列表
type ICMSQueryCityPost struct {
	CountryCode  string `json:"countryCode"`  // 国家代码
	RegionFirst  string `json:"regionFirst"`  // 地址一级区划名称（州/省）
	RegionSecond string `json:"regionSecond"` //地址二级区划名称（城市）
}

/*
{
    "countryCode":"CN",
    "regionFirst":"guangdong",
    "regionSecond":"Chaozhou"
}
*/

type ICMSQueryCityData struct {
	RegionSecondAbbr string `json:"regionSecondAbbr"` // 地址二级区划简称(城市)
	RegionSecond     string `json:"regionSecond"`     // 地址二级区划名称(城市)
	RegionSecondI18n string `json:"regionSecondI18n"` // 地址二级区划名称国际化(城市)
	ContinentCode    string `json:"continentCode"`    // 所属大洲
	CountryCode      string `json:"countryCode"`      // 国家编码
	CreateBy         string `json:"createBy"`         // 创建者账号
	CurrencyCode     string `json:"currencyCode"`     // 币种代码
	GmtCreate        string `json:"gmtCreate"`        // 创建时间
	GmtModified      string `json:"gmtModified"`      // 更新时间
	ID               int64  `json:"id"`               // ID
	ModifiedBy       string `json:"modifiedBy"`       // 更新者账号
	PhoneCode        string `json:"phoneCode"`        // 国际电话区号
	RegionFirst      string `json:"regionFirst"`      // 地址一级区划名称（州/省）
	SfCode           string `json:"sfCode"`           // 大网定义的ID（基本被弃用）
	Status           int    `json:"status"`           // 状态(0禁用，1启用)
	TimeZone         string `json:"timeZone"`         // 默认时区，如果一个国家跨多个时区，则通常为首都所在时区
	WeightUnit       string `json:"weightUnit"`       // 重量单位
}

type ICMSQueryCityResp struct {
	SFApiResponse
	Data []ICMSQueryCityData `json:"data"`
}

/*{
    "data": [
        {
            "continentCode": "AS",
            "countryCode": "CN",
            "createBy": "01405048",
            "gmtCreate": "2022-04-07T02:35:48.000+00:00",
            "gmtModified": "2022-11-03T02:20:49.000+00:00",
            "id": 233696,
            "modifiedBy": "01405048",
            "regionFirst": "Guangdong",
            "regionSecond": "Chaozhou",
            "regionSecondI18n": "Chaozhou",
            "status": 1
        }
    ],
    "success": true
}
*/

// 3.查询国家列表 ICMS_QUERY_COUNTRY_LIST

type ICMSCountryPost struct {
	ContinentCode string `json:"continentCode"` // 大洲代码
	CountryCode   string `json:"countryCode"`   // 国家代码
	PhoneCode     string `json:"phoneCode"`     // 手机区号
}

type ICMSCountryData struct {
	ID              int64  `json:"id"`              // ID
	ContinentCode   string `json:"continentCode"`   // 所属大洲
	CountryCode     string `json:"countryCode"`     // 国家二字码
	CountryName     string `json:"countryName"`     // 国家英文名称
	CountryNameI18n string `json:"countryNameI18n"` // 国家名称-指定语言
	CreateBy        string `json:"createBy"`        // 创建者账号
	CurrencyCode    string `json:"currencyCode"`    // 币种代码
	CustomsCode     string `json:"customsCode"`     // 中国海关对国家/地区的编码
	GmtCreate       string `json:"gmtCreate"`       // 创建时间
	GmtModified     string `json:"gmtModified"`     // 更新时间
	ModifiedBy      string `json:"modifiedBy"`      // 更新者账号
	NumberCode      string `json:"numberCode"`      // 国家数字码
	PhoneCode       string `json:"phoneCode"`       // 国际电话区号
	SfCode          string `json:"sfCode"`          // 大网定义的ID（基本被弃用）
	Status          int    `json:"status"`          // 状态(0禁用，1启用)
	ThreeCode       string `json:"threeCode"`       // 国家三字码
	TimeZone        string `json:"timeZone"`        // 默认时区
	WeightUnit      string `json:"weightUnit"`      // 重量单位
}

type ICMSCountryResp struct {
	SFApiResponse
	Data []ICMSCountryData `json:"data"`
}

// 4.地址区划-查询邮编属性 ICMS_QUERY_POST_CODE_ATTRIBUTE
type ICMSQueryCodeAttributePost struct{}

type ICMSQueryCodeAttributeData struct{}

type ICMSQueryCodeAttributeResp struct {
	SFApiResponse
	Data []ICMSQueryCodeAttributeData `json:"data"`
}

// 5.地址区划-根据国际区划查询邮编 ICMS_QUERY_POST_CODE_BY_REGION
type ICMSQueryCodeByRegionPost struct{}

type ICMSQueryCodeByRegionData struct{}

type ICMSQueryCodeByRegionResp struct {
	SFApiResponse
	Data []ICMSQueryCodeByRegionData `json:"data"`
}

// 6.地址区划-查询邮编详细信息 ICMS_QUERY_POST_CODE_INFO
type ICMSQueryCodeInfoPost struct{}

type ICMSQueryCodeInfoData struct{}

type ICMSQueryCodeInfoResp struct {
	SFApiResponse
	Data []ICMSQueryCodeInfoData `json:"data"`
}

// 7.地址区划-邮编模糊查询 ICMS_QUERY_POST_CODE_VAGUE
type ICMSQueryCodeVaguePost struct{}

type ICMSQueryCodeVagueData struct{}

type ICMSQueryCodeVagueResp struct {
	SFApiResponse
	Data []ICMSQueryCodeVagueData `json:"data"`
}

// 8.地址区划-查询州省列表 ICMS_QUERY_PROVINCE_LIST
type ICMSQueryProvinceListPost struct{}

type ICMSQueryProvinceListData struct{}

type ICMSQueryProvinceListResp struct {
	SFApiResponse
	Data []ICMSQueryProvinceListData `json:"data"`
}

// 9.地址区划-根据邮编查询国际区划 ICMS_QUERY_REGION_BY_POST_CODE
type ICMSQueryRegionByCodePost struct{}

type ICMSQueryRegionByCodeData struct{}

type ICMSQueryRegionByCodeResp struct {
	SFApiResponse
	Data []ICMSQueryRegionByCodeData `json:"data"`
}

// 10.地址区划-国际区划级联查询 ICMS_REGION_CASCADE_QUERY
type ICMSRegionCascadeQueryPost struct{}

type ICMSRegionCascadeQueryData struct{}

type ICMSRegionCascadeQueryResp struct {
	SFApiResponse
	Data []ICMSRegionCascadeQueryData `json:"data"`
}

// 11.地址区划-国际区划模糊查询 ICMS_REGION_VAGUE_QUERY
type ICMSRegionVagueQueryPost struct {
	CountryCode string `json:"countryCode"` // 国家代码
	Keyword     string `json:"keyword"`     // 关键词（必须为3个字符以上）
	PageNo      string `json:"pageNo"`      // 页码，默认1
	PageSize    string `json:"pageSize"`    // 每页数据量，默认100
}

type ICMSRegionVagueQueryData struct {
	RegionFirst       string `json:"regionFirst"`       // 地址一级区划名称（州/省）
	RegionFirstLocal  string `json:"regionFirstLocal"`  // 地址一级区划名称（州/省）-当地语言
	RegionFirstCn     string `json:"regionFirstCn"`     // 地址一级区划名称（州/省）-中文
	RegionSecond      string `json:"regionSecond"`      // 地址二级区划名称（城市）
	RegionSecondLocal string `json:"regionSecondLocal"` // 地址二级区划名称（城市）-当地语言
	RegionSecondCn    string `json:"regionSecondCn"`    // 地址二级区划名称（城市）-中文
	RegionThird       string `json:"regionThird"`       // 三级区划名称（区/县）
	RegionThirdLocal  string `json:"regionThirdLocal"`  // 三级区划名称（区/县）-当地语言
	RegionThirdCn     string `json:"regionThirdCn"`     // 三级区划名称（区/县）-中文
}

type ICMSRegionVagueQueryResp struct {
	SFApiResponse
	Data        []ICMSRegionVagueQueryData `json:"data"`
	PageNo      string                     `json:"pageNo"`
	PageSize    string                     `json:"pageSize"`
	Keyword     string                     `json:"keyword"`
	CountryCode string                     `json:"countryCode"`
}

// 12.地址区划-国际电话有效性校验 ICMS_PHONE_VALID_CHECK
type ICMSPhoneValidCheckPost struct {
	CountryCode                string `json:"countryCode"`                // 原寄地/目的地国家/地区二字码
	PhoneCode                  string `json:"phoneCode"`                  // 国际电话区号（纯数字）
	PhoneNumber                string `json:"phoneNumber"`                // 电话号码（纯数字）
	CheckMode                  string `json:"checkMode"`                  // 校验方式（默认值为0）
	DomesticPhoneCode          string `json:"domesticPhoneCode"`          // 国内电话区号
	ExtensionNumber            string `json:"extensionNumber"`            // 分机号
	PhoneCodeBelongCountryCode string `json:"phoneCodeBelongCountryCode"` // 区号归属国二字码
}

type ICMSPhoneValidCheckResp struct {
	SFApiResponse
	Result string `json:"result"`
}
