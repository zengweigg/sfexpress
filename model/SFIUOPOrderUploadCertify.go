package model

type SFIUOPOrderUploadCertify struct {
	CustomerCode    string `json:"customerCode"`
	CertCardNo      string `json:"certCardNo"`
	CertName        string `json:"certName"`
	SfWaybillNo     string `json:"sfWaybillNo"`
	CertType        string `json:"certType"`
	FileContent     string `json:"fileContent"`
	FileContentTwo  string `json:"fileContentTwo"`
	FileFormat      string `json:"fileFormat"`
	CustomerOrderNo string `json:"customerOrderNo"`
	Version         string `json:"version"`
}

type SFIUOPOrderUploadCertifyDeData struct {
	SFApiResponse
}

type SFIUOPOrderUploadCertifyResponse struct {
	SFResponse
	DeData SFIUOPOrderUploadCertifyDeData `json:"deData"`
}
