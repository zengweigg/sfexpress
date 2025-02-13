package tool

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"math/rand"
	"sort"
	"time"
)

const (
	ValidateSignatureError = -40001
	ParseXmlError          = -40002
	ComputeSignatureError  = -40003
	IllegalAesKey          = -40004
	ValidateAppidError     = -40005
	EncryptAESError        = -40006
	DecryptAESError        = -40007
	IllegalBuffer          = -40008
	EncodeBase64Error      = -40009
	DecodeBase64Error      = -40010
	GenReturnXmlError      = -40011
)

type BizMsgCrypt struct {
	aesKey []byte
	token  string
	appKey string
}

func NewBizMsgCrypt(token string, encodingAesKey string, appKey string) (*BizMsgCrypt, error) {
	if len(encodingAesKey) != 43 {
		return &BizMsgCrypt{nil, token, appKey}, errors.New(GetMessage(IllegalAesKey))
	}
	s, _ := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	return &BizMsgCrypt{s, token, appKey}, nil
}

func (biz *BizMsgCrypt) GetNetworkBytesOrder(sourceNumber int) []byte {
	orderBytes := make([]byte, 4)
	orderBytes[3] = (byte)(sourceNumber & 0xFF)
	orderBytes[2] = (byte)(sourceNumber >> 8 & 0xFF)
	orderBytes[1] = (byte)(sourceNumber >> 16 & 0xFF)
	orderBytes[0] = (byte)(sourceNumber >> 24 & 0xFF)
	return orderBytes
}

func (biz *BizMsgCrypt) RecoverNetworkBytesOrder(orderBytes []byte) int {
	sourceNumber := 0
	for i := 0; i < 4; i++ {
		sourceNumber <<= 8
		sourceNumber |= (int)(orderBytes[i] & 0xff)
	}
	return sourceNumber
}

func GetRandomStr() string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Int()
	s := ""
	for i := 0; i < 16; i++ {
		n := rand.Intn(len(base))
		s += string(base[n])
	}
	return s
}

func AppendByte(slice []byte, data []byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (biz *BizMsgCrypt) Encrypt(randomStr string, text string) (string, error) {
	randomStrBytes := GetBytes(randomStr)
	textBytes := GetBytes(text)
	textLenBytes := biz.GetNetworkBytesOrder(len(textBytes))
	appIdBytes := GetBytes(biz.appKey)
	bg := NewByteGroup()
	bg.byteArray = AppendByte(bg.byteArray, randomStrBytes) // [80, 108, 78, 70, 71, 100, 83, 67, 50, 119, 100, 56, 102, 50, 81, 110]
	bg.byteArray = AppendByte(bg.byteArray, textLenBytes)   // [0, 0, 4, 31]
	bg.byteArray = AppendByte(bg.byteArray, textBytes)      // [123, 34, 105, 110, 116, 101, 114, 80, 114, 111, 100, 117, 99, 116, 67, 111, 100, 101, 34, 58, 34, 73, 78, 84, 48, 50, 53, 53, 34, 44, 34, 112, 97, 114, 99, 101, 108, 81, 117, 97, 110, 116, 105, 116, 121, 34, 58, 34, 49, 34, 44, 34, 114, 101, 99, 101, 105, 118, 101, 114, 73, 110, 102, 111, 34, 58, 123, 34, 99, 111, 117, 110, 116, 114, 121, 34, 58, 34, 85, 83, 34, 44, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 50, 53, 54, 48, 32, 83, 32, +955 more]
	bg.byteArray = AppendByte(bg.byteArray, appIdBytes)     // [97, 98, 56, 99, 55, 98, 100, 51, 51, 55, 50, 55, 52, 97, 98, 51, 101, 57, 52, 53, 56, 98, 55, 100, 100, 48, 54, 55, 100, 100, 99, 50]
	padBytes := PKCS7Padding(len(bg.byteArray))
	bg.byteArray = AppendByte(bg.byteArray, padBytes)
	encrypted := CBCEncrypt(bg.byteArray, biz.aesKey)
	return encrypted, nil
}

// EncryptMsg 生成应答信息签名，replyMsg 应答信息，timeStamp 时间戳，nonce 随机串
func (biz *BizMsgCrypt) EncryptMsg(replyMsg string, timeStamp string, nonce string) Param {
	encrypt, err := biz.Encrypt(GetRandomStr(), replyMsg)
	if err != nil {
		return Param{}
	}
	if timeStamp == "" {
		timeStamp = fmt.Sprintf("%x", time.Now().UnixMilli())
	}
	signature := GetSHA256(biz.token, timeStamp, nonce, encrypt)
	return Param{encrypt, signature}
}

func (biz *BizMsgCrypt) DecryptMsg(postData string) (string, error) {
	str, err := biz.Decrypt(postData)
	if err != nil {
		return "", errors.New(GetMessage(DecryptAESError))
	}
	return str, nil
}

func (biz *BizMsgCrypt) Decrypt(text string) (string, error) {
	bytes, err := CBCDecrypt(text, biz.aesKey)
	if err != nil {
		return "", errors.New(GetMessage(DecryptAESError))
	}
	textLenBytes := bytes[16:20]
	textLen := biz.RecoverNetworkBytesOrder(textLenBytes)
	content := string(GetBytes(string(bytes[20 : 20+textLen])))
	fromAppKey := string(GetBytes(string(bytes[20+textLen : 52+textLen])))
	if fromAppKey != biz.appKey {
		return "", errors.New(GetMessage(ValidateAppidError))
	}
	return content, nil
}

func GetBytes(str string) []byte {
	b := []byte(str)
	br := bytes.NewReader(b)
	reader, _ := charset.NewReaderLabel("utf-8", br)
	b, _ = ioutil.ReadAll(reader)
	return b
}

func GetMessage(code int) string {
	switch code {
	case ValidateSignatureError:
		return "签名验证错误"
	case ParseXmlError:
		return "xml解析失败"
	case ComputeSignatureError:
		return "sha加密生成签名失败"
	case IllegalAesKey:
		return "SymmetricKey非法"
	case ValidateAppidError:
		return "appid校验失败"
	case EncryptAESError:
		return "aes加密失败"
	case DecryptAESError:
		return "aes解密失败"
	case IllegalBuffer:
		return "解密后得到的buffer非法"
	//case EncodeBase64Error:
	//	return "base64加密错误"
	//case DecodeBase64Error:
	//	return "base64解密错误"
	//case GenReturnXmlError:
	//	return "xml生成失败"
	default:
		return "" // cannot be
	}
}

func GetSHA256(token string, timestamp string, nonce string, encrypt string) string {
	ss := []string{token, timestamp, nonce, encrypt}
	sort.Strings(ss)
	str := ""
	for _, value := range ss {
		str += value
	}
	h := sha256.New()
	h.Write([]byte(str))
	res := h.Sum(nil)
	return fmt.Sprintf("%x", res)
}
