package model

type SFIECSOrderCancel struct {
	CancelReason    string `json:"cancelReason"`
	CustomerOrderNo string `json:"customerOrderNo"`
	SfWaybillNo     string `json:"sfWaybillNo"`
}

type SFIECSOrderCancelDeData struct {
	SFApiResponse
}

type SFIECSOrderCancelResponse struct {
	SFResponse
	DeData SFIECSOrderCancelDeData `json:"deData"`
}
