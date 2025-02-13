package model

// 查询城市列表

type SFICMSQueryCityList struct {
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

type SFICMSQueryCityData struct {
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

type SFICMSQueryCityResp struct {
	SFApiResponse
	Data []SFICMSQueryCityData `json:"data"`
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
