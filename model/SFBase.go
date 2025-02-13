package model

type SFBase struct {
	ApiResultCode int         `json:"apiResultCode"`
	ApiErrorMsg   string      `json:"apiErrorMsg"`
	ApiResultData interface{} `json:"apiResultData"`
	ApiTimestamp  int64       `json:"apiTimestamp"`
}
