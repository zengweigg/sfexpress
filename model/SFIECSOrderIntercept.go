package model

type SFIECSOrderIntercept struct {
	SfWaybillNo     string `json:"sfWaybillNo"`
	CustomerOrderNo string `json:"customerOrderNo"`
	InterceptType   string `json:"interceptType"`
	Remark          string `json:"remark"`
}

type SFIECSOrderInterceptDeData struct {
	SFApiResponse
}

type SFIECSOrderInterceptResponse struct {
	SFResponse
	DeData SFIECSOrderInterceptDeData `json:"deData"`
}
