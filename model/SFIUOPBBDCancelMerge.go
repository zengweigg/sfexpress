package model

type SFIUOPBBDCancelMerge struct {
	CustomerCode string   `json:"customerCode"`
	SfWayBillNos []string `json:"sfWayBillNos"`
}

type SFIUOPBBDCancelMergeDeData struct {
	SFApiResponse
}

type SFIUOPBBDCancelMergeResponse struct {
	SFResponse
	DeData SFIUOPBBDCancelMergeDeData `json:"deData"`
}
