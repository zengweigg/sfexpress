package model

type SFIUOPOrderRequireCertify struct {
	CustomerCode string `json:"customerCode"`
	RequestId    string `json:"requestId"`
	SfWaybillNo  string `json:"sfWaybillNo"`
}

type SFIUOPOrderRequireCertifyDeData struct {
	SFApiResponse
	Data struct {
		UploadedCertsInfoList []interface{} `json:"uploadedCertsInfoList"`
		CusCertsInfoList      []interface{} `json:"cusCertsInfoList"`
		Tips                  string        `json:"tips"`
	} `json:"data"`
}

type SFIUOPOrderRequireCertifyResponse struct {
	SFResponse
	DeData SFIUOPOrderRequireCertifyDeData `json:"deData"`
}
