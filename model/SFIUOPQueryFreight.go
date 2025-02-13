package model

type SFIUOPQueryFreight struct {
	CustomerCode     string `json:"customerCode"`
	InterProductCode string `json:"interProductCode"`
	ReceiverCountry  string `json:"receiverCountry"`
	ReceiverPostCode string `json:"receiverPostCode"`
	SenderCountry    string `json:"senderCountry"`
	SenderPostCode   string `json:"senderPostCode"`
	Weight           string `json:"weight"`
}

type SFIUOPQueryFreightDeData struct {
	SFApiResponse

	Data struct {
		Currency       string `json:"currency"`
		DiscountedRate string `json:"discountedRate"`
		TotalFreight   string `json:"totalFreight"`
		Weight         string `json:"weight"`
	} `json:"data"`
}

type SFIUOPQueryFreightResponse struct {
	SFResponse
	DeData SFIUOPQueryFreightDeData `json:"deData"`
}
