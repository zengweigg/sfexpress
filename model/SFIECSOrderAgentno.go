package model

type SFIECSOrderAgentno struct {
	Service     string             `json:"service"`
	WaybillList []SFIECSOrderAgent `json:"waybillList"`
}

type SFIECSOrderAgent struct {
	OrderNo     string `json:"orderNo"`
	SfWaybillNo string `json:"sfWaybillNo"`
}

type SFIECSOrderAgentnoSuccessData struct {
	AgentWaybillNo string `json:"agentWaybillNo"`
	Agentcode      string `json:"agentcode"`
	OrderNo        string `json:"orderNo"`
	SfWaybillNo    string `json:"sfWaybillNo"`
	AgentLabelUrl  string `json:"agentLabelUrl"`
}

type SFIECSOrderAgentnoSuccessDataWrongData struct {
	Description string `json:"description"`
	OrderNo     string `json:"orderNo"`
	SfWaybillNo string `json:"sfWaybillNo"`
	Status      string `json:"status"`
}

type SFIECSOrderAgentnoDeData struct {
	SFApiResponse
	Service     string                                   `json:"service"`
	SuccessData []SFIECSOrderAgentnoSuccessData          `json:"successData"`
	WrongData   []SFIECSOrderAgentnoSuccessDataWrongData `json:"wrongData"`
}

type SFIECSOrderAgentnoResponse struct {
	SFResponse
	DeData SFIECSOrderAgentnoDeData `json:"deData"`
}
