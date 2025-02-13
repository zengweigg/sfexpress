package model

type SFIUOPQueryTrack struct {
	CustomerCode string `json:"customerCode"`
	SfWaybillNo  string `json:"sfWaybillNo"`
}

type SFIUOPQueryTrackDeData struct {
	SFApiResponse
	Data []struct {
		Code             string `json:"code"`
		Msg              string `json:"msg"`
		SfWaybillNo      string `json:"sfWaybillNo"`
		TrackDetailItems []struct {
			Gmt               string `json:"gmt"`
			LocalTm           string `json:"localTm"`
			OpCode            string `json:"opCode"`
			ReasonCode        string `json:"reasonCode"`
			SfWaybillNo       string `json:"sfWaybillNo"`
			TrackAddr         string `json:"trackAddr"`
			TrackCountry      string `json:"trackCountry"`
			TrackOutRemark    string `json:"trackOutRemark"`
			TrackPostCode     string `json:"trackPostCode"`
			TrackRegionFirst  string `json:"trackRegionFirst"`
			TrackRegionSecond string `json:"trackRegionSecond"`
		} `json:"trackDetailItems"`
	} `json:"data"`
}

type SFIUOPQueryTrackResponse struct {
	SFResponse
	DeData SFIUOPQueryTrackDeData `json:"deData"`
}
