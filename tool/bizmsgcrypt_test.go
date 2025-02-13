package tool

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

func TestEncryptAndDecryptMsg(t *testing.T) {
	encodingAesKey := "eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dX"
	token := "auth_77860671-58ae-4d51-a6c4-11768db2c0d1_1644906436690"
	timestamp := "1637927849983"
	nonce := "1637927849983"
	appKey := "ab8c7bd337274ab3e9458b7dd067ddc2"
	replyMsg := "{\"interProductCode\":\"INT0255\",\"parcelQuantity\":\"1\",\"receiverInfo\":{\"country\":\"US\",\"address\":\"2560 S Barrington Ave Apt 101\",\"regionFirst\":\"California\",\"contact\":\"Lutz Olutola\",\"company\":\"Lutz Olutola\",\"postCode\":\"90064\",\"regionSecond\":\"Los Angeles\",\"email\":\"kalade@msn.com\",\"phoneNo\":\"3107775442\"},\"apiUsername\":\"popsandBox\",\"parcelInfoList\":[{\"unit\":\"个\",\"amount\":5,\"quantity\":1,\"hsCode\":\"1041748\",\"goodsUrl\":\"http://www.baidu.com\",\"cName\":\"白色定时烤暖鞋子烘干机\",\"name\":\"Electric Shoe Dryer\",\"invoiceNumber\":\"1234567890\",\"weight\":0.27,\"currency\":\"USD\"}],\"parcelTotalWeight\":\"0.1\",\"declaredCurrency\":\"USD\",\"senderInfo\":{\"country\":\"CN\",\"address\":\"YueQuan Dong RD.855\",\"regionFirst\":\"zhejiang\",\"contact\":\"Yuanyuan Zhou\",\"vat\":\"\",\"iossNo\":\"\",\"company\":\"Yuanyuan Zhou\",\"postCode\":\"322200\",\"regionSecond\":\"jinhua\",\"phoneNo\":\"18367985969\",\"eori\":\"\"},\"isBat\":\"0\",\"erpCode\":\"BINJIANG\",\"createOrderType\":\"1\",\"platformOrderId\":\"7438\",\"customerOrderNo\":\"SA1423285902376994\",\"platformCode\":\"prestashop\",\"paymentInfo\":{\"payMethod\":\"1\"},\"declaredValue\":5}"

	biz, err := NewBizMsgCrypt(token, encodingAesKey, appKey)
	if err != nil {
		t.Fail()
	}
	p := biz.EncryptMsg(replyMsg, timestamp, nonce)
	// go:   Param{encrypt='4NMNHJ2FwXP6VEkSGXIOkZiEl3sg5MPb+8uNyaXeaSUoxYsRfwJblmXRyjAi21BamPbwfg4uY99NL6LXsjBimLvaphN4XB4AQuhHnBxvZ0Rn0ScoryBM0jM193vLRxu64YBaEZfr+BsGSjLeSZHakcZR83dVc+CNSStwOF9Z16l/T6sNgExijRNILGBAJR0CUJ3N8Had3IZ1ScLp5P2iCkJuXOFr5stz/51nxeAOKcqTRlKO+A5swHr3Y5lp0Pn6gxomwB2hkEBUyids7wHNcoNFxYNwM7sV1if6BHvc3W6sQPKmRqpuoG8K7prlMgGRiPD2DF42mhLR2DIoDUJ+Q2+ahitKKZ81RBX7WA5s5SCWs5+cjAEgFLyU74e6pRwxumr5OtocqbUj/uyzO4lDE1ZDjgzJl6Z7d/I1ZUnrEayO6VT7qSK3s+ZWAgWRMBBzsFa/kRpBnwLkKNj0W2gK6UUYoqEe6FyNnUVS9L4qtoB+6OVIxCFl/jxAOi528ze5up8+eCJEPJBPxS+E8n8NLQqGvzPiaZbJN8wLV5yurwP8cVXy5nDlbP3RaLI5YOo/51m9ct/IdB5/Ln/gH+x6ut7rXcKHDAiOOJCy8gCsOgByRFIUQyemb68BHnpaAnrKBWQpfKjpQURcRYyEystsoQSzyyJmKdNbml75ymYcDmEEgWxbnvyvXl/4Q9JN126NVrwNhgY1VRYZKl0JhtdrJjMfiOKJzmhBD5FU33Ad8Lk9P/qYNGh8aQ8YITOWlKVYPCaFPTHS/hrbzIU3YrICEFHbgmn+kTgSsxQwlGR/aUCShFKYYgrsIT1ld2RFiRYTWgrg3YfsIMdZ7TTnbjPGz41xDaqnUygmUe7zjbe9ps72Jk2vY57KMZFFI9oBKFDef+v1ctBR5v8FesVpVlt2RFnn7OuCp1kq3hVvze30m4nZ7FQuNkJ4+H7nNeLbrNbp8uZPYvHoZms3fr+sAjsZwoamCuzxnNighd+FmT+cKS/oyZz30CacaB/DuPHW+kwFPpvoPAGUVfq21i2+eJ2OpMASwAFtiRKfULzAIL4eKT6rTLjY1uBLzF6u8QKCG06vIuEo7tSc9VzR/zlRW8m0n8faxLZyBsgfgg4tu54FemzqQWvNxagVkdyGrramVnpdTWKDWe2qFoLVklQBgGLgEwSCbbcfWFj9TIADKmVWKGOEqE2rVtR+98fX+RFLOLvsDvRGXVUc1jJccV1AJltxI1H3HJL9N/SeBwtNFDZmHZxiHaDsQjiay85mv3jRgIpGhzkkv+7VkEVtIv8k5uuy6rsHg0FLpGcNi7K6+qQmUrDq/9odlXQNwEPHo0Y7FA+99tJ6XkyTZGqThSjxPmv45ifN6SSJ8UD+K0LTRiTVFCiGvcKKJYFXw+Rl94t4TQWBYizFyTpWgeTEtXebsoz+CYP3xHOGqy5nIdb5DbOByhoMqEPEm9+hKldwi7r9OdJ+9kB7oFap/NKJmc0lgvkZRQ==', signature='2e4934dcdab2fe07b6ca4816f8b04ec99622e73da0e81d7b42930873980f544b'}
	// java: Param{encrypt='4NMNHJ2FwXP6VEkSGXIOkZiEl3sg5MPb+8uNyaXeaSUoxYsRfwJblmXRyjAi21BamPbwfg4uY99NL6LXsjBimLvaphN4XB4AQuhHnBxvZ0Rn0ScoryBM0jM193vLRxu64YBaEZfr+BsGSjLeSZHakcZR83dVc+CNSStwOF9Z16l/T6sNgExijRNILGBAJR0CUJ3N8Had3IZ1ScLp5P2iCkJuXOFr5stz/51nxeAOKcqTRlKO+A5swHr3Y5lp0Pn6gxomwB2hkEBUyids7wHNcoNFxYNwM7sV1if6BHvc3W6sQPKmRqpuoG8K7prlMgGRiPD2DF42mhLR2DIoDUJ+Q2+ahitKKZ81RBX7WA5s5SCWs5+cjAEgFLyU74e6pRwxumr5OtocqbUj/uyzO4lDE1ZDjgzJl6Z7d/I1ZUnrEayO6VT7qSK3s+ZWAgWRMBBzsFa/kRpBnwLkKNj0W2gK6UUYoqEe6FyNnUVS9L4qtoB+6OVIxCFl/jxAOi528ze5up8+eCJEPJBPxS+E8n8NLQqGvzPiaZbJN8wLV5yurwP8cVXy5nDlbP3RaLI5YOo/51m9ct/IdB5/Ln/gH+x6ut7rXcKHDAiOOJCy8gCsOgByRFIUQyemb68BHnpaAnrKBWQpfKjpQURcRYyEystsoQSzyyJmKdNbml75ymYcDmEEgWxbnvyvXl/4Q9JN126NVrwNhgY1VRYZKl0JhtdrJjMfiOKJzmhBD5FU33Ad8Lk9P/qYNGh8aQ8YITOWlKVYPCaFPTHS/hrbzIU3YrICEFHbgmn+kTgSsxQwlGR/aUCShFKYYgrsIT1ld2RFiRYTWgrg3YfsIMdZ7TTnbjPGz41xDaqnUygmUe7zjbe9ps72Jk2vY57KMZFFI9oBKFDef+v1ctBR5v8FesVpVlt2RFnn7OuCp1kq3hVvze30m4nZ7FQuNkJ4+H7nNeLbrNbp8uZPYvHoZms3fr+sAjsZwoamCuzxnNighd+FmT+cKS/oyZz30CacaB/DuPHW+kwFPpvoPAGUVfq21i2+eJ2OpMASwAFtiRKfULzAIL4eKT6rTLjY1uBLzF6u8QKCG06vIuEo7tSc9VzR/zlRW8m0n8faxLZyBsgfgg4tu54FemzqQWvNxagVkdyGrramVnpdTWKDWe2qFoLVklQBgGLgEwSCbbcfWFj9TIADKmVWKGOEqE2rVtR+98fX+RFLOLvsDvRGXVUc1jJccV1AJltxI1H3HJL9N/SeBwtNFDZmHZxiHaDsQjiay85mv3jRgIpGhzkkv+7VkEVtIv8k5uuy6rsHg0FLpGcNi7K6+qQmUrDq/9odlXQNwEPHo0Y7FA+99tJ6XkyTZGqThSjxPmv45ifN6SSJ8UD+K0LTRiTVFCiGvcKKJYFXw+Rl94t4TQWBYizFyTpWgeTEtXebsoz+CYP3xHOGqy5nIdb5DbOByhoMqEPEm9+hKldwi7r9OdJ+9kB7oFap/NKJmc0lgvkZRQ==', signature='2e4934dcdab2fe07b6ca4816f8b04ec99622e73da0e81d7b42930873980f544b'}
	fmt.Println(p.ToString())

	encrypt := "4NMNHJ2FwXP6VEkSGXIOkZiEl3sg5MPb+8uNyaXeaSUoxYsRfwJblmXRyjAi21BamPbwfg4uY99NL6LXsjBimLvaphN4XB4AQuhHnBxvZ0Rn0ScoryBM0jM193vLRxu64YBaEZfr+BsGSjLeSZHakcZR83dVc+CNSStwOF9Z16l/T6sNgExijRNILGBAJR0CUJ3N8Had3IZ1ScLp5P2iCkJuXOFr5stz/51nxeAOKcqTRlKO+A5swHr3Y5lp0Pn6gxomwB2hkEBUyids7wHNcoNFxYNwM7sV1if6BHvc3W6sQPKmRqpuoG8K7prlMgGRiPD2DF42mhLR2DIoDUJ+Q2+ahitKKZ81RBX7WA5s5SCWs5+cjAEgFLyU74e6pRwxumr5OtocqbUj/uyzO4lDE1ZDjgzJl6Z7d/I1ZUnrEayO6VT7qSK3s+ZWAgWRMBBzsFa/kRpBnwLkKNj0W2gK6UUYoqEe6FyNnUVS9L4qtoB+6OVIxCFl/jxAOi528ze5up8+eCJEPJBPxS+E8n8NLQqGvzPiaZbJN8wLV5yurwP8cVXy5nDlbP3RaLI5YOo/51m9ct/IdB5/Ln/gH+x6ut7rXcKHDAiOOJCy8gCsOgByRFIUQyemb68BHnpaAnrKBWQpfKjpQURcRYyEystsoQSzyyJmKdNbml75ymYcDmEEgWxbnvyvXl/4Q9JN126NVrwNhgY1VRYZKl0JhtdrJjMfiOKJzmhBD5FU33Ad8Lk9P/qYNGh8aQ8YITOWlKVYPCaFPTHS/hrbzIU3YrICEFHbgmn+kTgSsxQwlGR/aUCShFKYYgrsIT1ld2RFiRYTWgrg3YfsIMdZ7TTnbjPGz41xDaqnUygmUe7zjbe9ps72Jk2vY57KMZFFI9oBKFDef+v1ctBR5v8FesVpVlt2RFnn7OuCp1kq3hVvze30m4nZ7FQuNkJ4+H7nNeLbrNbp8uZPYvHoZms3fr+sAjsZwoamCuzxnNighd+FmT+cKS/oyZz30CacaB/DuPHW+kwFPpvoPAGUVfq21i2+eJ2OpMASwAFtiRKfULzAIL4eKT6rTLjY1uBLzF6u8QKCG06vIuEo7tSc9VzR/zlRW8m0n8faxLZyBsgfgg4tu54FemzqQWvNxagVkdyGrramVnpdTWKDWe2qFoLVklQBgGLgEwSCbbcfWFj9TIADKmVWKGOEqE2rVtR+98fX+RFLOLvsDvRGXVUc1jJccV1AJltxI1H3HJL9N/SeBwtNFDZmHZxiHaDsQjiay85mv3jRgIpGhzkkv+7VkEVtIv8k5uuy6rsHg0FLpGcNi7K6+qQmUrDq/9odlXQNwEPHo0Y7FA+99tJ6XkyTZGqThSjxPmv45ifN6SSJ8UD+K0LTRiTVFCiGvcKKJYFXw+Rl94t4TQWBYizFyTpWgeTEtXebsoz+CYP3xHOGqy5nIdb5DbOByhoMqEPEm9+hKldwi7r9OdJ+9kB7oFap/NKJmc0lgvkZRQ=="
	unEncrypt, err := biz.DecryptMsg(encrypt)
	if err != nil {
		errors.New(err.Error())
		t.Fail()
	}
	fmt.Println(unEncrypt)
}

func TestAppendByte(t *testing.T) {
	encodingAesKey := "eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dX"
	token := "auth_77860671-58ae-4d51-a6c4-11768db2c0d1_1644906436690"
	appKey := "ab8c7bd337274ab3e9458b7dd067ddc2"

	randomStr := "PlNFGdSC2wd8f2Qn"
	replyMsg := "{\"interProductCode\":\"INT0255\",\"parcelQuantity\":\"1\",\"receiverInfo\":{\"country\":\"US\",\"address\":\"2560 S Barrington Ave Apt 101\",\"regionFirst\":\"California\",\"contact\":\"Lutz Olutola\",\"company\":\"Lutz Olutola\",\"postCode\":\"90064\",\"regionSecond\":\"Los Angeles\",\"email\":\"kalade@msn.com\",\"phoneNo\":\"3107775442\"},\"apiUsername\":\"popsandBox\",\"parcelInfoList\":[{\"unit\":\"个\",\"amount\":5,\"quantity\":1,\"hsCode\":\"1041748\",\"goodsUrl\":\"http://www.baidu.com\",\"cName\":\"白色定时烤暖鞋子烘干机\",\"name\":\"Electric Shoe Dryer\",\"invoiceNumber\":\"1234567890\",\"weight\":0.27,\"currency\":\"USD\"}],\"parcelTotalWeight\":\"0.1\",\"declaredCurrency\":\"USD\",\"senderInfo\":{\"country\":\"CN\",\"address\":\"YueQuan Dong RD.855\",\"regionFirst\":\"zhejiang\",\"contact\":\"Yuanyuan Zhou\",\"vat\":\"\",\"iossNo\":\"\",\"company\":\"Yuanyuan Zhou\",\"postCode\":\"322200\",\"regionSecond\":\"jinhua\",\"phoneNo\":\"18367985969\",\"eori\":\"\"},\"isBat\":\"0\",\"erpCode\":\"BINJIANG\",\"createOrderType\":\"1\",\"platformOrderId\":\"7438\",\"customerOrderNo\":\"SA1423285902376994\",\"platformCode\":\"prestashop\",\"paymentInfo\":{\"payMethod\":\"1\"},\"declaredValue\":5}"
	textBytes := GetBytes(replyMsg)

	biz, err := NewBizMsgCrypt(token, encodingAesKey, appKey)
	if err != nil {
		t.Fail()
	}
	appIdBytes := GetBytes(biz.appKey)
	networkBytesOrder := biz.GetNetworkBytesOrder(len(textBytes))
	b := make([]byte, 1000)
	b = AppendByte(b, GetBytes(randomStr))
	b = AppendByte(b, networkBytesOrder)
	b = AppendByte(b, textBytes)
	b = AppendByte(b, appIdBytes)
	padBytes := PKCS7Padding(len(b))
	b = AppendByte(b, padBytes)
	fmt.Println(b)
}

func TestBase64(t *testing.T) {
	// java result: 120 75 -71 109 -10 -60 66 8 -39 119 118 -39 -20 -127 34 80 -52 -95 -40 -74 70 -87 104 107 -58 -95 73 -107 -108 -125 -27 -43
	//				120 75 185 109 246 196 66 8 217 119 118 217 236 129 34 80 204 161 216 182 70 169 104 107 198 161 73 149 148 131
	encodingAesKey := "eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dX"
	//				   eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dU=
	//				   eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dU=
	b, _ := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	fmt.Println(b)
	s := base64.StdEncoding.EncodeToString(b)
	fmt.Println(s)
}

func TestGetRandomStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetRandomStr())
	}
}

func TestGetSHA256(t *testing.T) {
	token := "auth_77860671-58ae-4d51-a6c4-11768db2c0d1_1644906436690"
	timestamp := "1637927849983"
	nonce := "1637927849983"
	encrypt := "J4f3NwHmlm9SH+v/foimkqptLyEOiPvXSygJ+7l7pg0bzxw0zV3xW6POtT4GDS81pU14nh/1RyKX/uKz4+2JX3g4XuJVFLQgSNiMfLcxmPSM+OOATUpNvovoMDeqARiZ5Gm1Iv2mJtl+N+QsntCEZ5rpwUaMMjXO9/xaajubnvq+NUAo/22igCsC0dPXtSiwuXNmtzEi3njE34pCIuoJzcJUq8l4UH9rYQsOBV988y5ZJxIovtl4Yoc4hk/WWoAsl+C293Q/UFQQda/s2HvBL0Df3faD+fXLp+AEc8XE+OpI/iW8a2lPEk5vEylbH7RhuB0aMBfzSVbNnaHQsOVad20gqsf9SawYovxp1R+p73TEabcz13GS5XpCFWOSX9CeVmTClXl2drimDyFIa129p0WJA18qo+MN/QWClmljRWgbNPaEF1R45fINdmHze0SryBT330ahWqgBg8ONIkSaRwebpqK7fNS7oj+sQheOJ/8ZdO/EM4/D6pSGdrVuEQcfXkgujRhgBzbYWqbS39AaUJQO10duf/eszrWOmxVg5W9Y+aLAqAJ+dJgymqiGUO8IHlCl8vXvas82VHhTF3vU9wui4WhpjGLpJ7bWYLCthbra00HX8yGdpULNm0cpf5+RFNy/bAFRmvTLcOhlk6EPBdBwhJkKZLTUIJPIx6K9OvDEI9kO0QVMxS+64xKBuxquI/Dz2v6aLAX+QZ6BrlxhtSguHfysODYV6E+lDKc6JoBTfB5jHktaT2xASE3yznxbtzKlu2mzKGSuqpMZu/kAQpl3/gjLbRnINR/i+878EoeXqjSvB+cLheSmUv12UF7cwznSXFwY1Ppguqs0wpC3FEfwqvPyjgtC5jGuIAh5xr4FKtznDPYMPgDjKqKemHPNfDO/zuxH/+yKB9KA5QCgG8s14Q/ieI6PdvheG0/xgFA16qlOPz8mfM127rjY5d1nzplYNyt5+8NGFCqT34OajotGZaDQlfq84jYQhGz52XcgiQCNWWimmi6y3tRZ2WPZJBQa0BWtZ6U9j6GnFSMSz2A0rT8AjDQAN+SwOrWHSJbuK0raes2L760CnyLNcDkYnmx0tzF30DcuUhMkL6+vxKouSk3XXHk1mN/avPTJnUM0+fCB+qUd6mD8O+i1H16A2PhqP2yM91pHl0MII88hzwais4n2tAOxEX5PdDJPICMfD2TqF5iiDmOWW8yqAwb4CCQA9sA7r1RZDrwwNKNnfzC4qj1ALriU7AIWmBKTLRAk5nowVmOf3EWr6kWF9UnpEkM45lk468RGT3P5EPd9HrCDa59TGuwS6BVTe7I/1GOgTykp8Mpzdvmcq7DJ39TdwQtIImX9ieJ2FRKb/xNGaUn2GEQ5lcwHCargImIczJoFzhLsAI/4ICxI0S7EQQMo0cvCd9n+kT7qoJKZJyXigwL0hi6RhjIV0oVrdTOjvyYCalLV5fBsOEtWwBNSQOuKNDDfi1MnFGMGJOtCqNfOAA=="
	// java 6d4783bb17cc966df325f4f251f60224ae40015cec8dc15d36f2690c2ba98437
	// go   6d4783bb17cc966df325f4f251f60224ae40015cec8dc15d36f2690c2ba98437
	s1 := GetSHA256(token, timestamp, nonce, encrypt)
	fmt.Println(s1)
}
