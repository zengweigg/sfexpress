package sfexpress

import (
	"fmt"
	"github.com/zengweigg/sfexpress/config"
	"github.com/zengweigg/sfexpress/model"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	sfClient := NewSFService(*cfg)

	// 构造测试请求数据
	//postData := model.ICMSQueryCityPost{
	//	CountryCode:  "CN",
	//	RegionFirst:  "guangdong",
	//	RegionSecond: "Chaozhou",
	//}

	// 测试获取城市列表
	//resp, err := sfClient.Services.Icms.GetCityList(postData)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println(resp.Data)

	// 构造测试请求数据
	//postData := model.ICMSRegionVagueQueryPost{
	//	CountryCode: "CN",
	//	Keyword:     "guangdong",
	//}

	//resp, err := sfClient.Services.Icms.RegionVagueQuery(postData)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println(resp.Data)
	postData := model.ICMSPhoneValidCheckPost{
		CountryCode: "CN",
		PhoneCode:   "86",
		PhoneNumber: "15125637820",
	}

	resp, err := sfClient.Services.Icms.PhoneValidCheck(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("666", resp.SFApiResponse.Code, resp.Result)

	// 测试空白
	//postData2 := model.SFBlank{}
	//resp2, err := sfClient.Services.Icms.GetContinentList(postData2)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println(resp2.Data)
}
