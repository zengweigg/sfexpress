package model

type SFIUOPOrderQuery struct {
	CustomerCode string `json:"customerCode"`
	SfWaybillNo  string `json:"sfWaybillNo"`
}

type SFIUOPOrderQueryDeData struct {
	SFApiResponse
	Data struct {
		AddServiceInfoList []struct {
			ServiceCode string `json:"serviceCode"`
			Value       string `json:"value"`
			ValueOne    string `json:"valueOne"`
		} `json:"addServiceInfoList"`
		AgentWaybillNo     string        `json:"agentWaybillNo"`
		ChildInfoList      []interface{} `json:"childInfoList"`
		CustomerOrderNo    string        `json:"customerOrderNo"`
		CustomerOrderNoTwo string        `json:"customerOrderNoTwo"`
		CustomsInfo        struct {
			AesNo               string `json:"aesNo"`
			BusinessRemark      string `json:"businessRemark"`
			CustomsBatch        string `json:"customsBatch"`
			ExportCustomsType   string `json:"exportCustomsType"`
			HarmonizedCode      string `json:"harmonizedCode"`
			ImportCustomsType   string `json:"importCustomsType"`
			SenderReasonContent string `json:"senderReasonContent"`
			TradeCondition      string `json:"tradeCondition"`
		} `json:"customsInfo"`
		DeclaredCurrency string  `json:"declaredCurrency"`
		DeclaredValue    float64 `json:"declaredValue"`
		InterProductCode string  `json:"interProductCode"`
		OrderExtendInfo  struct {
			IsSelfPick        string `json:"isSelfPick"`
			IsSignBack        string `json:"isSignBack"`
			SignBackWaybillNo string `json:"signBackWaybillNo"`
		} `json:"orderExtendInfo"`
		ParcelInfoList []struct {
			Amount           float64 `json:"amount"`
			Brand            string  `json:"brand"`
			EName            string  `json:"eName"`
			GoodsCode        string  `json:"goodsCode"`
			GoodsDesc        string  `json:"goodsDesc"`
			GoodsUrl         string  `json:"goodsUrl"`
			HsCode           string  `json:"hsCode"`
			Material         string  `json:"material"`
			Name             string  `json:"name"`
			OriginCountry    string  `json:"originCountry"`
			ProductCustomsNo string  `json:"productCustomsNo"`
			ProductRecordNo  string  `json:"productRecordNo"`
			Quantity         int     `json:"quantity"`
			Specifications   string  `json:"specifications"`
			StateBarCode     string  `json:"stateBarCode"`
			Unit             string  `json:"unit"`
		} `json:"parcelInfoList"`
		ParcelQuantity    int     `json:"parcelQuantity"`
		ParcelTotalWeight float64 `json:"parcelTotalWeight"`
		ParcelVolumeUnit  string  `json:"parcelVolumeUnit"`
		ParcelWeightUnit  string  `json:"parcelWeightUnit"`
		PaymentInfo       struct {
			PayMethod       string `json:"payMethod"`
			PayMonthCard    string `json:"payMonthCard"`
			TaxPayMethod    string `json:"taxPayMethod"`
			TaxPayMonthCard string `json:"taxPayMonthCard"`
		} `json:"paymentInfo"`
		PickupAppointTime     string `json:"pickupAppointTime"`
		PickupAppointTimeZone string `json:"pickupAppointTimeZone"`
		PickupType            string `json:"pickupType"`
		ReceiverInfo          struct {
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
			RegionFifth   string `json:"regionFifth"`
			RegionFirst   string `json:"regionFirst"`
			RegionSecond  string `json:"regionSecond"`
			RegionSixth   string `json:"regionSixth"`
			RegionThird   string `json:"regionThird"`
			TelAreaCode   string `json:"telAreaCode"`
			TelNo         string `json:"telNo"`
			Vat           string `json:"vat"`
		} `json:"receiverInfo"`
		Remark     string `json:"remark"`
		SenderInfo struct {
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
			RegionFifth   string `json:"regionFifth"`
			RegionFirst   string `json:"regionFirst"`
			RegionSecond  string `json:"regionSecond"`
			RegionSixth   string `json:"regionSixth"`
			RegionThird   string `json:"regionThird"`
			TelAreaCode   string `json:"telAreaCode"`
			TelNo         string `json:"telNo"`
			Vat           string `json:"vat"`
		} `json:"senderInfo"`
		SfWaybillNo string `json:"sfWaybillNo"`
	} `json:"data"`
}

type SFIUOPOrderQueryResponse struct {
	SFResponse
	DeData SFIUOPOrderQueryDeData `json:"deData"`
}
