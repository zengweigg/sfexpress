package model

type SFIUOPBBDPreMerge struct {
	Version      string   `json:"version"`
	CustomerCode string   `json:"customerCode"`
	SfWayBillNos []string `json:"sfWayBillNos"`
}

type SFIUOPBBDPreMergeDeData struct {
	SFApiResponse
	Data struct {
		PreMergeSfWaybillNo string `json:"preMergeSfWaybillNo"`
	} `json:"data"`
}

type SFIUOPBBDPreMergeResponse struct {
	SFResponse
	DeData SFIUOPBBDPreMergeDeData `json:"deData"`
}
