package model

type SFIECSOrder struct {
	ApiUsername       string       `json:"apiUsername"`
	CustomerOrderNo   string       `json:"customerOrderNo"`
	CustomsInfo       CustomsInfo  `json:"customsInfo"`
	DeclaredCurrency  string       `json:"declaredCurrency"`
	InterProductCode  string       `json:"interProductCode"`
	IsBat             bool         `json:"isBat"`
	ParcelInfoList    []ParcelItem `json:"parcelInfoList"`
	ParcelTotalWeight float64      `json:"parcelTotalWeight"`
	PaymentInfo       PaymentInfo  `json:"paymentInfo"`
	ReceiverInfo      ReceiverInfo `json:"receiverInfo"`
	SenderInfo        SenderInfo   `json:"senderInfo"`
}

type CustomsInfo struct {
	PassportId string `json:"passportId"`
	Category   string `json:"category"`
}

type ParcelItem struct {
	Amount        float64 `json:"amount"`
	CName         string  `json:"cName"`
	Currency      string  `json:"currency"`
	HsCode        string  `json:"hsCode"`
	Name          string  `json:"name"`
	Quantity      int     `json:"quantity"`
	Unit          string  `json:"unit"`
	Weight        float64 `json:"weight"`
	GoodsUrl      string  `json:"goodsUrl"`
	InvoiceNumber string  `json:"invoiceNumber"`
}

type PaymentInfo struct {
	PayMethod    string `json:"payMethod"`
	PayMonthCard string `json:"payMonthCard"`
}

type ReceiverInfo struct {
	Address      string `json:"address"`
	RegionSecond string `json:"regionSecond"`
	RegionThird  string `json:"regionThird"`
	Company      string `json:"company"`
	Contact      string `json:"contact"`
	Country      string `json:"country"`
	Email        string `json:"email"`
	PhoneNo      string `json:"phoneNo"`
	PostCode     string `json:"postCode"`
	RegionFirst  string `json:"regionFirst"`
	TelNo        string `json:"telNo"`
}

type SenderInfo struct {
	Address      string `json:"address"`
	RegionSecond string `json:"regionSecond"`
	RegionThird  string `json:"regionThird"`
	Company      string `json:"company"`
	Contact      string `json:"contact"`
	Country      string `json:"country"`
	PhoneNo      string `json:"phoneNo"`
	PostCode     string `json:"postCode"`
	RegionFirst  string `json:"regionFirst"`
	TelNo        string `json:"telNo"`
}

type SFIECSOrderResponse struct {
	SFResponse
	DeData SFIECSOrderDeData `json:"deData"`
}

type SFIECSOrderDeData struct {
	SFApiResponse
	Data struct {
		AgentCode       string `json:"agentCode"`
		CustomerOrderNo string `json:"customerOrderNo"`
		DirectionCode   string `json:"directionCode"`
		SfWaybillNo     string `json:"sfWaybillNo"`
	} `json:"data"`
}
