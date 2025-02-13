package config

type Config struct {
	APIKey         string
	APISecret      string
	EncodingAesKey string // APP encodingAesKey
	CustomerCode   string // APP 客户编码
	ApiUsername    string // APP IECS出口电商业务客户名称
	MonthCard      string // 月结卡号
	Lang           string // 语言
	Debug          bool   // 是否启用调试模式
	Sandbox        bool   // 是否为沙箱环境
	Timeout        int    // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:         "cdd169d87a4698d314754a79f180308a",
		APISecret:      "69fb3e868a45d0bdd5b0de0ea1db6ece",
		EncodingAesKey: "oI1YdU1cb1YC70HVjZRa3wXBLUsrIUYr5lDh0gFrMbe",
		CustomerCode:   "ICRME000SRN93",
		ApiUsername:    "popsandBox",
		MonthCard:      "9999999999",
		Lang:           "zh-CN",
		Debug:          true,
		Sandbox:        true,
		Timeout:        360,
	}
}
