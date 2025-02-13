package model

type SFIUOPBBDMergeResult struct {
	CustomerCode     string `json:"customerCode"`
	MergeSfWaybillNo string `json:"mergeSfWaybillNo"`
	SfWaybillNo      string `json:"sfWaybillNo"`
}

type SFIUOPBBDMergeResultDeData struct {
	SFApiResponse
	Data []struct {
		MergeRelationList []struct {
			MergeSfWaybillNo string `json:"mergeSfWaybillNo"`
			SfWaybillNos     string `json:"sfWaybillNos"`
		} `json:"mergeRelationList"`
		MergeSfWaybillNos string `json:"mergeSfWaybillNos"`
	} `json:"data"`
}

type SFIUOPBBDMergeResultResponse struct {
	SFResponse
	DeData SFIUOPBBDMergeResultDeData `json:"deData"`
}
