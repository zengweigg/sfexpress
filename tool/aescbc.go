package tool

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// CBCEncrypt AES CBC加密，text待加密明文，key密钥
func CBCEncrypt(text []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	iv := make([]byte, 16)
	copy(iv, key[0:16])
	blockMode := cipher.NewCBCEncrypter(block, iv)
	result := make([]byte, len(text))
	blockMode.CryptBlocks(result, text)
	return base64.StdEncoding.EncodeToString(result)
}

// CBCDecrypt AES CBC解密，encrypt 待解密的密文，key 密钥
func CBCDecrypt(encrypt string, key []byte) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(encrypt)
	iv := make([]byte, 16)
	copy(iv, key[0:16])
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(data))
	blockMode.CryptBlocks(result, data)
	return result, nil
}
