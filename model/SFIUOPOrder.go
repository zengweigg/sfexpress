package model

type SFIUOPOrder struct {
	ApiUsername           string                     `json:"apiUsername"`
	CustomerCode          string                     `json:"customerCode"`
	DeclaredCurrency      string                     `json:"declaredCurrency"`
	DeclaredValue         int                        `json:"declaredValue"`
	OrderOperateType      string                     `json:"orderOperateType"`
	ParcelQuantity        int                        `json:"parcelQuantity"`
	ParcelTotalWeight     float64                    `json:"parcelTotalWeight"`
	ParcelWeightUnit      string                     `json:"parcelWeightUnit"`
	ParcelVolumeUnit      string                     `json:"parcelVolumeUnit"`
	ParcelTotalLength     string                     `json:"parcelTotalLength"`
	ParcelTotalWidth      string                     `json:"parcelTotalWidth"`
	ParcelTotalHeight     string                     `json:"parcelTotalHeight"`
	PickupAppointTime     string                     `json:"pickupAppointTime"`
	PickupAppointTimeZone string                     `json:"pickupAppointTimeZone"`
	PickupType            string                     `json:"pickupType"`
	Remark                string                     `json:"remark"`
	SfWaybillNo           string                     `json:"sfWaybillNo"`
	Version               string                     `json:"version"`
	CustomerOrderNo       string                     `json:"customerOrderNo"`
	CustomerOrderNoTwo    string                     `json:"customerOrderNoTwo"`
	InterProductCode      string                     `json:"interProductCode"`
	IsBbd                 int                        `json:"isBbd"`
	AgentWaybillNo        string                     `json:"agentWaybillNo"`
	PaymentInfo           SFIUOPOrderPaymentInfo     `json:"paymentInfo"`
	AddServiceInfoList    []SFIUOPOrderService       `json:"addServiceInfoList"`
	ChildInfoList         []SFIUOPOrderChild         `json:"childInfoList"`
	CustomsInfo           SFIUOPOrderCustomsInfo     `json:"customsInfo"`
	OrderExtendInfo       SFIUOPOrderOrderExtendInfo `json:"orderExtendInfo"`
	ParcelInfoList        []SFIUOPOrderParcel        `json:"parcelInfoList"`
	ReceiverInfo          SFIUOPOrderAddress         `json:"receiverInfo"`
	SenderInfo            SFIUOPOrderAddress         `json:"senderInfo"`
	EcommerceInfo         SFIUOPOrderEcommerceInfo   `json:"ecommerceInfo"`
	IsBat                 bool                       `json:"isBat"`
}

type SFIUOPOrderPaymentInfo struct {
	PayMethod       string `json:"payMethod"`
	PayMonthCard    string `json:"payMonthCard"`
	TaxPayMethod    string `json:"taxPayMethod"`
	TaxPayMonthCard string `json:"taxPayMonthCard"`
}

type SFIUOPOrderService struct {
	ServiceCode string `json:"serviceCode"`
	Value       string `json:"value"`
	ValueOne    string `json:"valueOne"`
}

type SFIUOPOrderChild struct {
	SfWaybillNo string `json:"sfWaybillNo"`
	Length      string `json:"length"`
	Width       string `json:"width"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
}

type SFIUOPOrderCustomsInfo struct {
	AesNo               string `json:"aesNo"`
	BusinessRemark      string `json:"businessRemark"`
	CustomsBatch        string `json:"customsBatch"`
	HarmonizedCode      string `json:"harmonizedCode"`
	SenderReasonContent string `json:"senderReasonContent"`
	TradeCondition      string `json:"tradeCondition"`
}

type SFIUOPOrderOrderExtendInfo struct {
	IsSelfPick        string `json:"isSelfPick"`
	IsSignBack        string `json:"isSignBack"`
	SignBackWaybillNo string `json:"signBackWaybillNo"`
	BusinessMode      string `json:"businessMode"`
	SenderDeptCode    string `json:"senderDeptCode"`
	OperEmpCode       string `json:"operEmpCode"`
}

type SFIUOPOrderParcel struct {
	Amount           float64 `json:"amount"`
	Brand            string  `json:"brand"`
	Currency         string  `json:"currency"`
	GoodsCode        string  `json:"goodsCode"`
	GoodsDesc        string  `json:"goodsDesc"`
	GoodsUrl         string  `json:"goodsUrl"`
	HsCode           string  `json:"hsCode"`
	Name             string  `json:"name"`
	OriginCountry    string  `json:"originCountry"`
	ProductCustomsNo string  `json:"productCustomsNo"`
	ProductRecordNo  string  `json:"productRecordNo"`
	Quantity         int     `json:"quantity"`
	StateBarCode     string  `json:"stateBarCode"`
	Unit             string  `json:"unit"`
}

type SFIUOPOrderAddress struct {
	Address       string `json:"address"`
	CargoType     int    `json:"cargoType"`
	CertCardNo    string `json:"certCardNo"`
	CertType      string `json:"certType"`
	Company       string `json:"company"`
	Contact       string `json:"contact"`
	Country       string `json:"country"`
	Email         string `json:"email"`
	Eori          string `json:"eori"`
	PhoneAreaCode string `json:"phoneAreaCode"`
	PhoneNo       string `json:"phoneNo"`
	PostCode      string `json:"postCode"`
	TelAreaCode   string `json:"telAreaCode"`
	TelNo         string `json:"telNo"`
	Vat           string `json:"vat"`
	RegionFirst   string `json:"regionFirst"`
	RegionSecond  string `json:"regionSecond"`
	RegionThird   string `json:"regionThird"`
}

type SFIUOPOrderEcommerceInfo struct {
	ActPay                   int    `json:"actPay"`
	ApplicationNo            string `json:"applicationNo"`
	BuyerEmail               string `json:"buyerEmail"`
	BuyerFreight             int    `json:"buyerFreight"`
	BuyerNickname            string `json:"buyerNickname"`
	BuyerPremium             int    `json:"buyerPremium"`
	Currency                 string `json:"currency"`
	EnterpriseCheckCode      string `json:"enterpriseCheckCode"`
	EnterpriseCrossCode      string `json:"enterpriseCrossCode"`
	EnterpriseCustomsCode    string `json:"enterpriseCustomsCode"`
	EnterpriseName           string `json:"enterpriseName"`
	PayEnterpriseCheckCode   string `json:"payEnterpriseCheckCode"`
	PayEnterpriseCustomsCode string `json:"payEnterpriseCustomsCode"`
	PayNo                    string `json:"payNo"`
	PayTime                  string `json:"payTime"`
	PayToolType              string `json:"payToolType"`
	PayerCertNo              string `json:"payerCertNo"`
	PayerCertType            string `json:"payerCertType"`
	PayerName                string `json:"payerName"`
	PayerTel                 string `json:"payerTel"`
	PlatformCheckCode        string `json:"platformCheckCode"`
	PlatformCustomsCode      string `json:"platformCustomsCode"`
	PlatformDomainName       string `json:"platformDomainName"`
	PlatformName             string `json:"platformName"`
	PreferentialAmount       int    `json:"preferentialAmount"`
	PromotionType            string `json:"promotionType"`
	SellerAlipayNo           string `json:"sellerAlipayNo"`
	SellerEmail              string `json:"sellerEmail"`
	SellerNickname           string `json:"sellerNickname"`
	ShopWebCode              string `json:"shopWebCode"`
	StoreCode                string `json:"storeCode"`
	Tax                      string `json:"tax"`
}

type SFIUOPOrderData struct {
	IdentityUploadUrl  string   `json:"identityUploadUrl"`
	ChildWaybillNoList []string `json:"childWaybillNoList"`
	InvoiceUrl         string   `json:"invoiceUrl"`
	LabelUrl           string   `json:"labelUrl"`
	CustomerOrderNo    string   `json:"customerOrderNo"`
	SfWaybillNo        string   `json:"sfWaybillNo"`
}

type SFIUOPOrderDeData struct {
	SFApiResponse
	Data SFIUOPOrderData `json:"data"`
}

type SFIUOPOrderResponse struct {
	SFResponse
	DeData SFIUOPOrderDeData `json:"deData"`
}
