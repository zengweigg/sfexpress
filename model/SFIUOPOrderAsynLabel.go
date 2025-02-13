package model

type SFIUOPOrderAsynLabel struct {
	CustomerCode          string                           `json:"customerCode"`
	LabelModeSize         int                              `json:"labelModeSize"`
	PrintType             string                           `json:"printType"`
	PrintWaybillNoDtoList []SFIUOPOrderLabelPrintWaybillNo `json:"printWaybillNoDtoList"`
}

type SFIUOPOrderAsynLabelDeData struct {
	SFApiResponse
	Data struct {
		PrintRequestId string `json:"printRequestId"`
	} `json:"data"`
}

type SFIUOPOrderAsynLabelResponse struct {
	SFResponse
	DeData SFIUOPOrderAsynLabelDeData `json:"deData"`
}
