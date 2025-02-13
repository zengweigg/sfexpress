package model

type SFIUOPOrderAsynPrintResult struct {
	CustomerCode   string `json:"customerCode"`
	PrintRequestId string `json:"printRequestId"`
}

type SFIUOPOrderAsynPrintResultDeData struct {
	SFApiResponse
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

type SFIUOPOrderAsynPrintResultResponse struct {
	SFResponse
	DeData SFIUOPOrderAsynPrintResultDeData `json:"deData"`
}
