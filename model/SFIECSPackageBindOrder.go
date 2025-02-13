package model

type SFIECSPackageBindOrder struct {
	ApiUsername  string                     `json:"apiUsername"`
	PackageCode  string                     `json:"packageCode"`
	PackageCode2 string                     `json:"packageCode2"`
	PackageCode3 string                     `json:"packageCode3"`
	WaybillList  []SFIECSPackageBindOrderNo `json:"waybillList"`
}

type SFIECSPackageBindOrderNo struct {
	SfWaybillNo    string `json:"sfWaybillNo"`
	AgentWaybillNo string `json:"agentWaybillNo"`
	Weight         string `json:"weight"`
}

type SFIECSPackageBindOrderDeData struct {
	SFApiResponse
	Data SFIECSPackageBindOrderData `json:"data"`
}

type SFIECSPackageBindOrderData struct {
	PackageCode  string                     `json:"packageCode"`
	PackageCode2 string                     `json:"packageCode2"`
	PackageCode3 string                     `json:"packageCode3"`
	WaybillList  []SFIECSPackageBindOrderNo `json:"waybillList"`
}

type SFIECSPackageBindOrderResponse struct {
	SFResponse
	DeData SFIECSPackageBindOrderDeData `json:"deData"`
}
