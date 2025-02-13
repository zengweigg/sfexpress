package model

type SFIECSOrderLabel struct {
	PrintWaybillNoDtoList []SFIECSLabelPrintWaybillNo `json:"printWaybillNoDtoList"`
	ApiUsername           string                      `json:"apiUsername"`
	PrintPicking          bool                        `json:"printPicking"`
}

type SFIECSLabelPrintWaybillNo struct {
	SfWaybillNo string `json:"sfWaybillNo"`
}

type SFIECSOrderLabelResponse struct {
	SFResponse
	DeData SFIECSOrderLabelDeData `json:"deData"`
}

type SFIECSOrderLabelDeData struct {
	SFApiResponse
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}
