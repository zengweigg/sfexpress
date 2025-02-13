package sfexpress

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hiscaler/gox/filex"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type Token struct {
	AccessToken     string `json:"accessToken"`
	ExpiresIn       int    `json:"expireIn"`
	ExpiresDatetime int64  `json:"expires_datetime"`
}

// Valid 判断是否有token且没有过期
func (t Token) Valid() bool {
	return t.AccessToken != "" && t.ExpiresDatetime > time.Now().Unix()
}

type FileToken struct {
	Key   string
	Path  string
	Token Token
}

type TokenWriterReader interface {
	Read(key string) (Token, error)
	Write(key string, token Token) (bool, error)
	GetValidAccessToken(key string, s authorizationService) (Token, error)
}

// tokenFilePath 生成并返回存储Token的文件路径。
// 参数 key 用于区分不同的Token文件，通常是与Token相关联的唯一标识。
// 返回的路径是操作系统临时目录下的一个文件路径，文件名为 "sf_token_<key>.json"。
func tokenFilePath(key string) string {
	return path.Join(os.TempDir(), fmt.Sprintf("sf_token_%s.json", key))
}

// Read 从文件中读取与指定 key 关联的 Token 信息。
// 参数 key 用于标识唯一的 Token 文件。
// 返回 Token 结构体以及可能的错误信息。
func (ft FileToken) Read(key string) (Token, error) {
	token := Token{} // 初始化一个空的 Token 结构体
	var err error

	// 获取与 key 关联的 Token 文件路径
	file := tokenFilePath(key)

	// 检查文件是否存在
	if filex.Exists(file) {
		var b []byte
		// 读取文件内容
		if b, err = ioutil.ReadFile(file); err == nil {
			// 将文件内容解析为 Token 结构体
			err = json.Unmarshal(b, &token)
		}
	} else {
		// 如果文件不存在，返回错误
		err = errors.New("token file is not exists")
	}

	// 返回读取到的 Token 和可能的错误
	return token, err
}

// Write 将 Token 结构体序列化为 JSON 并写入与指定 key 关联的文件中。
// 参数 key 用于标识唯一的 Token 文件，token 是需要存储的 Token 数据。
// 返回一个布尔值表示写入是否成功，以及可能的错误信息。
func (ft FileToken) Write(key string, token Token) (bool, error) {
	// 将 Token 结构体序列化为 JSON 格式的字节数组
	b, err := json.Marshal(token)
	if err != nil {
		return false, err // 如果序列化失败，返回错误
	}

	// 将 JSON 数据写入文件，文件路径由 tokenFilePath(key) 生成，权限设置为 0777（可读可写可执行）
	err = ioutil.WriteFile(tokenFilePath(key), b, 0777)
	if err != nil {
		return false, err // 如果写入文件失败，返回错误
	}

	// 写入成功，返回 true 和 nil 错误
	return true, nil
}

// GetValidAccessToken 获取有效的 AccessToken
func (ft FileToken) GetValidAccessToken(key string, s authorizationService) (Token, error) {
	// 尝试从本地读取 Token
	token, err := ft.Read(key)
	if err == nil && token.Valid() {
		return token, nil
	}

	// 如果本地 Token 无效或不存在，则从远程获取新的 Token
	newToken, err := s.GetToken()
	if err != nil {
		return Token{}, err
	}

	// 将新的 Token 写入本地文件
	_, err = ft.Write(key, newToken)
	if err != nil {
		return Token{}, err
	}

	return newToken, nil
}
