package model

type SFIUOPOrderLabel struct {
	CustomerCode          string                           `json:"customerCode"`
	LabelModeSize         int                              `json:"labelModeSize"`
	PrintType             string                           `json:"printType"`
	PrintWaybillNoDtoList []SFIUOPOrderLabelPrintWaybillNo `json:"printWaybillNoDtoList"`
}

type SFIUOPOrderLabelPrintWaybillNo struct {
	IsPrintSubParent int    `json:"isPrintSubParent"`
	SfWaybillNo      string `json:"sfWaybillNo"`
}

type SFIUOPOrderLabelDeData struct {
	SFApiResponse
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

type SFIUOPOrderLabelResponse struct {
	SFResponse
	DeData SFIUOPOrderLabelDeData `json:"deData"`
}
