package sfexpress

import (
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/sfexpress/config"
	"github.com/zengweigg/sfexpress/tool"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	Iecs iecsService
	Iuop iuopService
	Ioms iomsService
	Ades adesservice
	Gts  gtsService
	Icms icmsService
}

// GetSign 消息加密获取签名
func GetSign(param, token, aeskey, apikey, timestamp string) (tool.Param, error) {
	biz, err := tool.NewBizMsgCrypt(token, aeskey, apikey)
	if err != nil {
		return tool.Param{}, err
	}
	p := biz.EncryptMsg(param, timestamp, timestamp)
	return p, nil
}

// DecodeMsg 解密消息
func DecodeMsg(response, token, aeskey, apikey string) (string, error) {
	biz, err := tool.NewBizMsgCrypt(token, aeskey, apikey)
	if err != nil {
		return "", err
	}
	return biz.DecryptMsg(response)
}
