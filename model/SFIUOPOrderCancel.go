package model

type SFIUOPOrderCancel struct {
	CustomerCode string `json:"customerCode"`
	SfWaybillNo  string `json:"sfWaybillNo"`
}

type SFIUOPOrderCancelDeData struct {
	SFApiResponse
}

type SFIUOPOrderCancelResponse struct {
	SFResponse
	DeData SFIUOPOrderCancelDeData `json:"deData"`
}
