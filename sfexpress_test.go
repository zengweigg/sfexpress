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
	postData := model.SFICMSQueryCityList{
		CountryCode:  "CN",
		RegionFirst:  "guangdong",
		RegionSecond: "Chaozhou",
	}
	// 测试获取城市列表
	resp, err := sfClient.Services.Icms.GetCityList(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(resp.Data)
}
