package model

type IUOPValidityCertify struct {
	CardNo   string `json:"cardNo"`
	TelPhone string `json:"telPhone"`
	UserName string `json:"userName"`
}

type IUOPValidityCertifyDeData struct {
	SFApiResponse
	Data bool `json:"data"`
}

type IUOPValidityCertifyResponse struct {
	SFResponse
	DeData IUOPValidityCertifyDeData `json:"deData"`
}
