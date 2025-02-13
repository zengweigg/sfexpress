package tool

import (
	"errors"
	"fmt"
	"testing"
)

func TestCBCEncrypt(t *testing.T) {
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
}

func TestCBCDecrypt(t *testing.T) {
	encodingAesKey := "eEu5bfbEQgjZd3bZ7IEiUMyh2LZGqWhrxqFJlZSD5dX"
	token := "auth_77860671-58ae-4d51-a6c4-11768db2c0d1_1644906436690"
	appKey := "ab8c7bd337274ab3e9458b7dd067ddc2"
	biz, err := NewBizMsgCrypt(token, encodingAesKey, appKey)
	if err != nil {
		t.Fail()
	}
	encrypt := "4NMNHJ2FwXP6VEkSGXIOkZiEl3sg5MPb+8uNyaXeaSUoxYsRfwJblmXRyjAi21BamPbwfg4uY99NL6LXsjBimLvaphN4XB4AQuhHnBxvZ0Rn0ScoryBM0jM193vLRxu64YBaEZfr+BsGSjLeSZHakcZR83dVc+CNSStwOF9Z16l/T6sNgExijRNILGBAJR0CUJ3N8Had3IZ1ScLp5P2iCkJuXOFr5stz/51nxeAOKcqTRlKO+A5swHr3Y5lp0Pn6gxomwB2hkEBUyids7wHNcoNFxYNwM7sV1if6BHvc3W6sQPKmRqpuoG8K7prlMgGRiPD2DF42mhLR2DIoDUJ+Q2+ahitKKZ81RBX7WA5s5SCWs5+cjAEgFLyU74e6pRwxumr5OtocqbUj/uyzO4lDE1ZDjgzJl6Z7d/I1ZUnrEayO6VT7qSK3s+ZWAgWRMBBzsFa/kRpBnwLkKNj0W2gK6UUYoqEe6FyNnUVS9L4qtoB+6OVIxCFl/jxAOi528ze5up8+eCJEPJBPxS+E8n8NLQqGvzPiaZbJN8wLV5yurwP8cVXy5nDlbP3RaLI5YOo/51m9ct/IdB5/Ln/gH+x6ut7rXcKHDAiOOJCy8gCsOgByRFIUQyemb68BHnpaAnrKBWQpfKjpQURcRYyEystsoQSzyyJmKdNbml75ymYcDmEEgWxbnvyvXl/4Q9JN126NVrwNhgY1VRYZKl0JhtdrJjMfiOKJzmhBD5FU33Ad8Lk9P/qYNGh8aQ8YITOWlKVYPCaFPTHS/hrbzIU3YrICEFHbgmn+kTgSsxQwlGR/aUCShFKYYgrsIT1ld2RFiRYTWgrg3YfsIMdZ7TTnbjPGz41xDaqnUygmUe7zjbe9ps72Jk2vY57KMZFFI9oBKFDef+v1ctBR5v8FesVpVlt2RFnn7OuCp1kq3hVvze30m4nZ7FQuNkJ4+H7nNeLbrNbp8uZPYvHoZms3fr+sAjsZwoamCuzxnNighd+FmT+cKS/oyZz30CacaB/DuPHW+kwFPpvoPAGUVfq21i2+eJ2OpMASwAFtiRKfULzAIL4eKT6rTLjY1uBLzF6u8QKCG06vIuEo7tSc9VzR/zlRW8m0n8faxLZyBsgfgg4tu54FemzqQWvNxagVkdyGrramVnpdTWKDWe2qFoLVklQBgGLgEwSCbbcfWFj9TIADKmVWKGOEqE2rVtR+98fX+RFLOLvsDvRGXVUc1jJccV1AJltxI1H3HJL9N/SeBwtNFDZmHZxiHaDsQjiay85mv3jRgIpGhzkkv+7VkEVtIv8k5uuy6rsHg0FLpGcNi7K6+qQmUrDq/9odlXQNwEPHo0Y7FA+99tJ6XkyTZGqThSjxPmv45ifN6SSJ8UD+K0LTRiTVFCiGvcKKJYFXw+Rl94t4TQWBYizFyTpWgeTEtXebsoz+CYP3xHOGqy5nIdb5DbOByhoMqEPEm9+hKldwi7r9OdJ+9kB7oFap/NKJmc0lgvkZRQ=="
	unEncrypt, err := biz.DecryptMsg(encrypt)
	if err != nil {
		errors.New(err.Error())
		t.Fail()
	}
	fmt.Println(unEncrypt)
}
