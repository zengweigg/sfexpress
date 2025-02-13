package model

type SFTokenData struct {
	ExpireIn    int    `json:"expireIn"`
	AccessToken string `json:"accessToken"`
}

type SFToken struct {
	SFBase
	ApiResultData SFTokenData `json:"apiResultData"`
}
