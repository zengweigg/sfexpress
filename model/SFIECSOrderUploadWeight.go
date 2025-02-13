package model

type SFIECSOrderUploadWeight struct {
	WaybillList []SFIECSOrderWeight `json:"waybillList"`
}

type SFIECSOrderWeight struct {
	Weight      string `json:"weight"`
	SfWaybillNo string `json:"sfWaybillNo"`
}

type SFIECSOrderUploadWeightDeData struct {
	SFApiResponse
}

type SFIECSOrderUploadWeightResponse struct {
	SFResponse
	DeData SFIECSOrderUploadWeightDeData `json:"deData"`
}
