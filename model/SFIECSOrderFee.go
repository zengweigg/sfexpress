package model

type SFIECSOrderFee struct {
	WeightAndFreightList []SFIECSweightAndFreightNo `json:"weightAndFreightList"`
}

type SFIECSweightAndFreightNo struct {
	SfWaybillNo    string `json:"sfWaybillNo"`
	AgentWaybillNo int    `json:"agentWaybillNo"`
}

type SFIECSOrderFeeDeData struct {
	SFApiResponse
	Data []struct {
		AgentWaybillNo   string  `json:"agentWaybillNo"`
		Clearance        float64 `json:"clearance"`
		ClearanceService float64 `json:"clearanceService"`
		DebitTime        int64   `json:"debitTime"`
		Freight          float64 `json:"freight"`
		Message          string  `json:"message"`
		RealWeight       float64 `json:"realWeight"`
		SfWaybillNo      string  `json:"sfWaybillNo"`
		Tariff           float64 `json:"tariff"`
		Weight           float64 `json:"weight"`
		WeightTime       int64   `json:"weightTime"`
	} `json:"data"`
}

type SFIECSOrderFeeResponse struct {
	SFResponse
	DeData SFIECSOrderFeeDeData `json:"deData"`
}
