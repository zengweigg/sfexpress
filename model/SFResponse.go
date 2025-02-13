package model

type SFResponse struct {
	ApiResultCode int    `json:"apiResultCode"`
	ApiResultData string `json:"apiResultData"`
	ApiTimestamp  int64  `json:"apiTimestamp"`
	ApiErrorMsg   string `json:"apiErrorMsg"`
}

type SFApiResponse struct {
	Msg     string `json:"msg"`
	Code    string `json:"code"`
	Success bool   `json:"success"`
}
