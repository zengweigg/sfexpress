package model

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

// 通用查询

// 1.路由查询 GTS_QUERY_TRACK
type GtsQueryTrackPost struct {
	SfWaybillNoList []string `json:"sfWaybillNoList"` // 顺丰运单号，一次查询最多支持10个单号
	PhoneNo         string   `json:"phoneNo"`         // 手机号后四位（下单时使用的手机号mobile）
}

type TrackInfo struct {
	OpCode            string `json:"opCode"`            // 操作码
	ReasonCode        string `json:"reasonCode"`        // SF国际原因代码
	SfWaybillNo       string `json:"sfWaybillNo"`       // 顺丰运单号
	TrackCountry      string `json:"trackCountry"`      // 轨迹发生国家
	TrackRegionFirst  string `json:"trackRegionFirst"`  // 轨迹发生一级区划 州/省
	TrackRegionSecond string `json:"trackRegionSecond"` // 轨迹发生二级区划 城市
	TrackAddr         string `json:"trackAddr"`         // 轨迹发生详细地址
	TrackPostCode     string `json:"trackPostCode"`     // 轨迹发生邮编
	LocalTm           string `json:"localTm"`           // 轨迹发生时间（当地时间），格式: yyyy-MM-dd HH:mm:ss
	Gmt               string `json:"gmt"`               // 标准时区格式：GMT+/- 08:00 or GMT+/-08:30
	TrackOutRemark    string `json:"trackOutRemark"`    // 路由轨迹对外描述
}

type GtsQueryTrackData struct {
	SfWaybillNo      string      `json:"sfWaybillNo"`      // 1. 顺丰单号
	Code             null.String `json:"code"`             // 2. 否
	Msg              null.String `json:"msg"`              // 3. 否
	TrackDetailItems []TrackInfo `json:"trackDetailItems"` // 4. 海外代理单号
}

type GtsQueryTrackResp struct {
	SFApiResponse
	Data []GtsQueryTrackData `json:"data"`
}

// 2.单票派送证明查询 IOMS_SEND_PROVE_ONEFILE
type IomsSendProveOnefilePost struct {
	CustomerCode string `json:"customerCode"` // 客户编码
	MonthlyCard  string `json:"monthlyCard"`  // 月结账号
	WaybillNo    string `json:"waybillNo"`    // 顺丰运单号
	FileType     string `json:"fileType"`     // 文件类型说明： 15：派送证明  71：微派任务
}

type IomsSendProveOnefileData struct {
	WaybillNo       string `json:"waybillNo"`       // 顺丰运单号
	WaybillFileName string `json:"waybillFileName"` // 顺丰运单号下的文件名
	FileBase64      string `json:"fileBase64"`      // 文件的Base64字符串
}

type IomsSendProveOnefileResp struct {
	SFApiResponse
	Data struct {
		CustomerCode    string                     `json:"customerCode"`    // 客户编码
		MonthlyCard     string                     `json:"monthlyCard"`     // 月结账号
		TaskCode        string                     `json:"taskCode"`        // 文件生成任务编码
		ResultCode      string                     `json:"resultCode"`      // 响应code
		ResultMsg       string                     `json:"resultMsg"`       // 响应msg
		Status          int                        `json:"status"`          // 文件生成任务处理状态 0：待处理 1：处理成功 2：处理失败 3：取消
		TaskBeginDate   int64                      `json:"taskBeginDate"`   // 任务有效期开始时间
		TaskEndDate     int64                      `json:"taskEndDate"`     // 任务有效期结束时间
		TaskCreateTime  int64                      `json:"taskCreateTime"`  // 文件生成任务创建时间
		HandleTime      int64                      `json:"handleTime"`      // 文件生成任务处理完成时间
		WaybillFileList []IomsSendProveOnefileData `json:"waybillFileList"` // 返回的运单查询结果
	} `json:"data"`
}

// 3.批量派送证明查询 IOMS_SEND_PROVE_CREATE
type IomsSendProveCreatePost struct {
	CustomerCode string   `json:"customerCode"` // 客户编码
	MonthlyCard  string   `json:"monthlyCard"`  // 月结账号
	WaybillList  []string `json:"waybillList"`  // 运单号集合（一次最多50条）
	FileType     string   `json:"fileType"`     // 文件类型（派送证明传15即可） 15：派送证明
}

type IomsSendProveCreateResp struct {
	SFApiResponse
	Data struct {
		CustomerCode  string   `json:"customerCode"`  // 客户编码
		MonthlyCard   string   `json:"monthlyCard"`   // 月结账号
		WaybillList   []string `json:"waybillList"`   // 运单号集合
		FileType      string   `json:"fileType"`      // 文件类型
		ResultCode    string   `json:"resultCode"`    // 响应原因码
		ResultMsg     string   `json:"resultMsg"`     // 响应原因描述
		TaskCode      string   `json:"taskCode"`      // 文件生成任务编码
		TaskBeginDate int64    `json:"taskBeginDate"` // 任务有效期开始时间
		TaskEndDate   int64    `json:"taskEndDate"`   // 任务有效期结束时间
	} `json:"data"`
}

// 第2步 查询

type IomsSendProveQueryPost struct {
	CustomerCode string `json:"customerCode"` // 客户编码
	MonthlyCard  string `json:"monthlyCard"`  // 月结账号
	TaskCode     string `json:"taskCode"`     // 运单文件查询服务返回的文件生成任务编码
}

type WaybillFile struct {
	WaybillNo       string `json:"waybillNo"`       // 顺丰运单号
	WaybillFileName string `json:"waybillFileName"` // 顺丰运单号下的文件名
	FileBase64      string `json:"fileBase64"`      // 文件的Base64字符串
}

type WaybillInfo struct {
	WaybillNo       string        `json:"waybillNo"`       // 顺丰运单号
	WaybillFileList []WaybillFile `json:"waybillFileList"` // 顺丰运单号下的文件集合
}

type IomsSendProveQueryResp struct {
	SFApiResponse
	Data struct {
		CustomerCode    string        `json:"customerCode"`    // 客户编码
		MonthlyCard     string        `json:"monthlyCard"`     // 月结账号
		TaskCode        string        `json:"taskCode"`        // 文件生成任务编码
		ResultCode      string        `json:"resultCode"`      // 响应code
		ResultMsg       string        `json:"resultMsg"`       // 响应msg
		Status          int           `json:"status"`          // 文件生成任务处理状态 0：待处理 1：处理成功 2：处理失败 3：取消
		TaskBeginDate   time.Time     `json:"taskBeginDate"`   // 任务有效期开始时间
		TaskEndDate     time.Time     `json:"taskEndDate"`     // 任务有效期结束时间
		TaskCreateTime  time.Time     `json:"taskCreateTime"`  // 文件生成任务创建时间
		HandleTime      time.Time     `json:"handleTime"`      // 文件生成任务处理完成时间
		WaybillInfoList []WaybillInfo `json:"waybillInfoList"` // 返回的运单查询结果
	} `json:"data"`
}

// 4.计费重和费用项查询 IOMS_QUERY_CHARGE
type IomsQueryChargePost struct {
	CustomerCode string   `json:"customerCode"` // 客户编码
	MonthlyCard  string   `json:"monthlyCard"`  // 月结账号
	WaybillList  []string `json:"waybillList"`  // 顺丰运单号集合（最多10条）
}

type WaybillFee struct {
	FeeId              string  `json:"feeId"`              // 费用类型ID
	FeeTypeCode        string  `json:"feeTypeCode"`        // 费用类型代码
	FeeAmt             float64 `json:"feeAmt"`             // 费用金额
	DestCurrencyCode   string  `json:"destCurrencyCode"`   // 费用币种
	PaymentTypeCode    string  `json:"paymentTypeCode"`    // 费用付款方式，（收方、寄方、第三方）
	SettlementTypeCode string  `json:"settlementTypeCode"` // 结算方式（现结、月结）
	TicketOffsetAmt    float64 `json:"ticketOffsetAmt"`    // 券抵免金额
	DeductionAmt       float64 `json:"deductionAmt"`       // 抵免金额
}

type WaybillCharge struct {
	WaybillNo         string       `json:"waybillNo"`         // 顺丰运单号
	ChildWaybillNos   []string     `json:"childWaybillNos"`   // 子单号(包裹号)List
	RealWeightQty     float64      `json:"realWeightQty"`     // 实际重
	Volume            float64      `json:"volume"`            // 体积
	MeterageWeightQty float64      `json:"meterageWeightQty"` // 计费重
	Quantity          int          `json:"quantity"`          // 包裹件数
	WaybillFeeList    []WaybillFee `json:"waybillFeeList"`    // 费用明细集合
}

type IomsQueryChargeResp struct {
	SFApiResponse
	Data struct {
		CustomerCode      string          `json:"customerCode"`      // 客户编码
		MonthlyCard       string          `json:"monthlyCard"`       // 月结账号
		ResultCode        string          `json:"resultCode"`        // 接口响应CODE
		ResultMsg         string          `json:"resultMsg"`         // 接口响应描述
		WaybillChargeList []WaybillCharge `json:"waybillChargeList"` // 返回的运单费用信息集合
	} `json:"data"`
}

// 5.尾程供应商派送证明查询 ADES_AGENT_QUERY_POD
type AdesAgentQueryPodPost struct {
	SfWaybillNo string `json:"sfWaybillNo"` // 顺丰运单号
}

type AdesAgentQueryPodResp struct {
	SFApiResponse
	Data struct {
		SfWaybillNo    string `json:"sfWaybillNo"`    // 顺丰运单号
		AgentWaybillNo string `json:"agentWaybillNo"` // 物流商代理单号
		FileBase64     string `json:"fileBase64"`     // PDF文件的Base64字符串
		PdfHtmlStr     string `json:"pdfHtmlStr"`     // 原始HTML信息
	} `json:"data"`
}

// 6.获取税金信息及支付渠道 IOMS_TAX_PAYMENT
type IomsTaxPaymentPost struct {
	BNo   string `json:"BNo"`   // 运单号
	Phone string `json:"phone"` // 收方或者寄方手机号
}

type IomsTaxPaymentData struct{}

type FreightInfo struct {
	TotalFreightFee float64 `json:"totalFreightFee"` // 总运费
	TotalFee        float64 `json:"totalFee"`        // 总费用
	FreightFee      float64 `json:"freightFee"`      // 运费
	Currency        string  `json:"currency"`        // 费用币种
	WaybillNo       string  `json:"waybillNo"`       // 运单号
	ValueAddedFee   float64 `json:"valueAddedFee"`   // 增值服务费
}

type TaxInfo struct {
	HandlingFee        float64     `json:"handlingFee"`        // 报关手续费
	Interdna           string      `json:"interdna"`           // 清关口岸对应的国家（国别地区）
	TaxAccount         string      `json:"taxAccount"`         // 税金结算账号（寄付、第三方必填）
	DecCode            string      `json:"decCode"`            // 报关单号
	Currency           string      `json:"currency"`           // 费用币种
	CustomsDate        string      `json:"customsDate"`        // 实际报关日期
	DutyValue          float64     `json:"dutyValue"`          // 货物税
	TotalValue         float64     `json:"totalValue"`         // 税金合计
	WareHouseFee       float64     `json:"wareHouseFee"`       // 报关仓储费
	MainbNo            string      `json:"mainbNo"`            // 海关提供提单号
	IsAudit            null.String `json:"isAudit"`            // 已审核:Y；未审核：N
	OtherFee           float64     `json:"otherFee"`           // 其他费用
	GstValue           float64     `json:"gstValue"`           // 消费税
	InDate             string      `json:"inDate"`             // 货物到达口岸日期
	IncrementFee       float64     `json:"incrementFee"`       // 增值税
	DecValue           float64     `json:"decValue"`           // 申报价值
	Mawb               string      `json:"mawb"`               // 航空主单号
	TaxPerNo           string      `json:"taxPerNo"`           // 纳税人编号
	OtherDutyAmount    float64     `json:"otherDutyAmount"`    // 其他税
	CusDeptCode        string      `json:"cusDeptCode"`        // 口岸三字码，如LAX
	TotalFee           float64     `json:"totalFee"`           // 税费合计
	QuarantineFee      float64     `json:"quarantineFee"`      // 检验检疫费
	TaxPaymentType     string      `json:"taxPaymentType"`     // 税金付款方式（1：寄方月结，2：到方现结，3：第三方支付，4：其他）
	CusCurrencySymbol  string      `json:"cusCurrencySymbol"`  // 申报币别
	TraiffValue        float64     `json:"traiffValue"`        // 关税
	CusNetCode         string      `json:"cusNetCode"`         // 口岸公司核算代码，如755Z（模板配置）
	DecType            string      `json:"decType"`            // 报关类型
	TotalFeeVal        float64     `json:"totalFeeVal"`        // 税费合计
	WaybillNo          string      `json:"waybillNo"`          // 运单号
	TaxPaymentTypeName string      `json:"taxPaymentTypeName"` // 税金付款方式名称
}

// FeeInfo 定义运单费用信息
type FeeInfo struct {
	TotalFee float64     `json:"totalFee"` // 总费用
	Freight  FreightInfo `json:"freight"`  // 运费项
	Currency string      `json:"currency"` // 费用币种
	Tax      TaxInfo     `json:"tax"`      // 税金项
}

type PaymentCompany struct {
	Code           string `json:"code"`           // 编码
	PayChannelName string `json:"payChannelName"` // 渠道名称
	PayChannelLogo string `json:"payChannelLogo"` // 渠道logo base64
	Name           string `json:"name"`           // 名称
	Logo           string `json:"logo"`           // base64
	PayChannel     string `json:"payChannel"`     // 渠道
}

type IomsTaxPaymentResp struct {
	SFApiResponse
	Data struct {
		CargoPieces          float64          `json:"cargoPieces"`          // 物品数量
		Interdna             string           `json:"interdna"`             // 清关口岸对应的国家（国别地区），如US
		ConsignorCityCode    string           `json:"consignorCityCode"`    // 寄件城市编码
		PaymentCompanies     []PaymentCompany `json:"paymentCompanies"`     // 支付公司列表
		ConsignorCountryCode string           `json:"consignorCountryCode"` // 寄件国编码,如US
		Fee                  FeeInfo          `json:"fee"`                  // 费用信息
		AddresseeCountryCode string           `json:"addresseeCountryCode"` // 目的国编码,如US
		SrcCompany           string           `json:"srcCompany"`           // 寄件公司名称
		EnglishCargoNames    string           `json:"englishCargoNames"`    // 英文物品名称
		DecCode              string           `json:"decCode"`              // 报关单号
		PhoneStatus          bool             `json:"phoneStatus"`          // 电话号码状态，true-通过、false-未通过
		Currency             string           `json:"currency"`             // 口岸税金币种，如USD
		DesPhone             string           `json:"desPhone"`             // 收件电话
		SrcPhone             string           `json:"srcPhone"`             // 寄件电话
		CustomsDate          string           `json:"customsDate"`          // 实际报关日期
		DesAddress           string           `json:"desAddress"`           // 收件地址
		InDate               string           `json:"inDate"`               // 货物到达口岸日期
		CusDeptCode          string           `json:"cusDeptCode"`          // 口岸三字码，如LAX
		TaxPerNo             string           `json:"taxPerNo"`             // 纳税人编号
		CargoNames           string           `json:"cargoNames"`           // 物品名称
		CargoWeight          float64          `json:"cargoWeight"`          // 物品重量
		CusNetCode           string           `json:"cusNetCode"`           // 口岸公司核算代码，如755Z
		DecType              string           `json:"decType"`              // 报关类型
		Status               int              `json:"status"`               // 税金支付状态（1:已支付、2:待支付、3:异常、4:已关闭）
		StatusDetail         int              `json:"statusDetail"`         // 支付状态详情
		CusCurrencySymbol    string           `json:"cusCurrencySymbol"`    // 申报币别
		DesCompany           string           `json:"desCompany"`           // 收件公司名称
		DesPhoneCode         string           `json:"desPhoneCode"`         // 收件电话区号
		SrcPhoneCode         string           `json:"srcPhoneCode"`         // 寄件电话区号
		Bno                  string           `json:"bno"`                  // 运单号
		MainbNo              string           `json:"mainbNo"`              // 海关提供提单号
		SrcPersonName        string           `json:"srcPersonName"`        // 寄件联系人
		DesPersonName        string           `json:"desPersonName"`        // 收件联系人
		Mawb                 string           `json:"mawb"`                 // 航空主单号
	} `json:"data"`
}

// 7.获取税金支付链接 IOMS_TAX_PAYMENT_LINK
type IomsTaxPaymentLinkPost struct {
	BNo                string `json:"BNo"`                // 运单号
	PaymentCompanyCode string `json:"paymentCompanyCode"` // 支付渠道/公司代码
}

type IomsTaxPaymentLinkResp struct {
	SFApiResponse
	Data struct {
		QrUrl   string `json:"qrUrl"`   // 二维码地址
		PayType int    `json:"payType"` // 支付方式 1-二维码 2-跳转
		PayUrl  string `json:"payUrl"`  // 支付链接-跳转
	} `json:"data"`
}

// 8.获取在线支付短链接 IOMS_PAYMENT_LINK
type IomsPaymentLinkPost struct {
	BNo   string `json:"BNo"`   // 运单号
	Payer string `json:"payer"` // 支付方 1:收方,2:寄方
}

type IomsPaymentLinkResp struct {
	SFApiResponse
	Data struct {
		WaybillNo string   `json:"waybillNo"` // 主运单号
		PayURL    string   `json:"payURL"`    // 支付链接
		ValidTime int64    `json:"validTime"` // 链接有效时间-单位：分钟
		SubList   []string `json:"subList"`   // 子运单号集合
	} `json:"data"`
}

// 9.派送证明上传 ADES_UPLOAD_POD
type AdesUploadPodPost struct {
	SfWaybillNo    string      `json:"sfWaybillNo"`         // 顺丰运单号
	FileContent    string      `json:"fileContent"`         // Base64格式的pod文件,fileType=URL时该字段传URL的值，该值必须为可以直接GET访问的文件地址
	FileFormat     string      `json:"fileFormat"`          // 文件类型，支持PDF/PNG/JPG/URL
	AgentCode      string      `json:"agentCode"`           // 代理商代码
	AgentWaybillNo string      `json:"agentWaybillNo"`      // 代理商单号
	Sku            null.String `json:"sku,omitempty"`       // SKU
	Quantity       null.Int    `json:"quantity,omitempty"`  // 签收数量
	Timestamp      null.Int    `json:"timestamp,omitempty"` // 签收时间毫秒时间戳
}

type AdesUploadPodResp struct {
	SFApiResponse
	Timestamp int64 `json:"time"`
}

// 10.操作运单费用明细查询 IOMS_WAYBILL_BILL
type IomsWaybillBillPost struct {
	WaybillNo string `json:"waybillNo"` // 顺丰运单号
}

// WaybillBillFeeDto 定义费用信息
type WaybillBillFeeDto struct {
	FeeTypeCode        string  `json:"feeTypeCode"`        // 费用类型
	FeeAmt             float64 `json:"feeAmt"`             // 折前金额
	CurrencyCode       string  `json:"currencyCode"`       // 币别代码
	PaymentTypeCode    string  `json:"paymentTypeCode"`    // 付款方式 1.寄付 2.到付 3.第三方付
	SettlementTypeCode string  `json:"settlementTypeCode"` // 结算方式 1.现结 2.月结
	TicketOffsetAmt    float64 `json:"ticketOffsetAmt"`    // 券抵免金额
	DeductionAmt       float64 `json:"deductionAmt"`       // 抵免金额
}

type IomsWaybillBillResp struct {
	SFApiResponse
	Data struct {
		WaybillNo          string              `json:"waybillNo"`          // 运单号
		RealWeightQty      float64             `json:"realWeightQty"`      // 实际重量
		Volume             float64             `json:"volume"`             // 体积
		MeterageWeightQty  float64             `json:"meterageWeightQty"`  // 计费重
		UnitWeight         string              `json:"unitWeight"`         // 重量单位
		Quantity           float64             `json:"quantity"`           // 包裹数量
		Length             float64             `json:"length"`             // 长度
		Height             float64             `json:"height"`             // 高度
		Width              float64             `json:"width"`              // 宽度
		WaybillBillFeeList []WaybillBillFeeDto `json:"waybillBillFeeList"` // 费用信息列表
	} `json:"data"`
}

// 11.中国海关税金信息查询 ICSM_CN_TAX_QUERY
type ICSMCnTaxQueryPost struct {
	BNo         string      `json:"BNo"`             // 运单号
	MonthlyCard string      `json:"monthlyCard"`     // 运月结卡号
	Phone       null.String `json:"phone,omitempty"` // 收方或者寄方手机号
}

type ICSMCnTaxQueryResp struct {
	SFApiResponse
	Data struct {
		Bno             string  `json:"bno"`             // 运单号
		CusDeptCode     string  `json:"cusDeptCode"`     // 口岸代码
		DecCode         string  `json:"decCode"`         // 报关单编码/报关单号
		Currency        string  `json:"currency"`        // 税费币种,如CNY
		TaxAccount      string  `json:"taxAccount"`      // 税金结算账号
		TraiffValue     float64 `json:"traiffValue"`     // 关税
		IncrementFee    float64 `json:"incrementFee"`    // 增值税
		GstValue        float64 `json:"gstValue"`        // 消费税
		DutyValue       float64 `json:"dutyValue"`       // 货物税
		OtherDutyAmount float64 `json:"otherDutyAmount"` // 其他税金
		QuarantineFee   float64 `json:"quarantineFee"`   // 检验检疫费
		CommissionFee   float64 `json:"commissionFee"`   // 委托外贸代理费
		OtherFee        float64 `json:"otherFee"`        // 其它费用
		HandlingFee     float64 `json:"handlingFee"`     // 报关手续费
		WareHouseFee    float64 `json:"wareHouseFee"`    // 报关仓储费
		TotalValue      float64 `json:"totalValue"`      // 税金合计
		TotalFee        float64 `json:"totalFee"`        // 费用合计
		TotalFeeVal     float64 `json:"totalFeeVal"`     // 税费合计
	} `json:"data"`
}

// 12.代理商更新运单 ADES_AGENT_UPDATE_WAYBILL
type AdesAgentUpdateWaybillPost struct {
	FromSourceSys string       `json:"fromSourceSys"` // 来源系统编码-每个代理给固定值(请联系对接人进行申请)
	Action        string       `json:"action"`        // 固定值：agentUpdateOrder
	ProxyCode     string       `json:"proxyCode"`     // 代理标识-每个代理给固定值(请联系对接人进行申请)
	SfRefNo       string       `json:"sfRefNo"`       // 顺丰运单号
	Version       string       `json:"version"`       // 版本号固定传输v2.0
	Uuid          string       `json:"uuid"`          // 请求唯一标识-可用后续查询
	CreateTime    int64        `json:"createTime"`    // 请求数据的时间戳
	Pdata         TrackingInfo `json:"pdata"`         // 数据实体
}

type TrackingInfo struct {
	TrackingNo string `json:"trackingNo"` // 代理运单号
	Carrier    string `json:"carrier"`    // 承运商
	WaybillNo  string `json:"waybillNo"`  // 顺丰运单号
}

// 13.国际收寄标准 SOP_COLLECT_SHIP_STANDARD
type SopCollectShipStandardPost struct {
	RequestId           string      `json:"requestId"`           // 请求id
	OperateType         string      `json:"operateType"`         // 固定值1，代表下单
	CallTm              int64       `json:"callTm"`              // 当前接口调用时间戳
	SrcCountryAreaCode  string      `json:"srcCountryAreaCode"`  // （始发地）原寄地国家/地区代码二字码 （港澳台则设置为3位数字代码，如852表示香港；CN表示大陆）
	DestCountryAreaCode string      `json:"destCountryAreaCode"` // 目的地国家/地区代码二字码（港澳台则设置为3位数字代码，如852），如852表示香港；CN表示大陆
	ConsignorRemark     string      `json:"consignorRemark"`     // 寄方件种属性代码（01：个人件 02：公司件）
	AddresseeRemark     string      `json:"addresseeRemark"`     // 收方件种属性代码（01：个人件 02：公司件）
	Lang                null.String `json:"lang,omitempty"`      // 国际化语言代码，如：简体中文：zh-CN，繁体中文：zh-TW，英文：en
	GdescDtos           []GdescDto  `json:"gdescDtos"`           // （托寄物）品名信息集合
}

type GdescDto struct {
	Quantity null.String `json:"quantity,omitempty"` // 数量
	Unit     null.String `json:"unit,omitempty"`     // 单位
	Gdesc    string      `json:"gdesc"`              // 物品名称 ：删除数字【0-9】和特殊字符后不能为空，【特殊字符】。纯中文长度必须大于等于2，纯英文必须大于2
	Price    null.String `json:"price,omitempty"`    // 单价
	Weight   null.String `json:"weight,omitempty"`   // 单托寄物重量
	Label    null.String `json:"label,omitempty"`    // 物品标签（例如：易燃易爆）
}

type SopCollectShipStandardData struct {
	Gdesc   string `json:"gdesc"`   // 托寄物名称（品名）
	IsFuzzy string `json:"isFuzzy"` // 是否模糊匹配（Y:是 N:否）
}

type SopCollectShipStandardResp struct {
	SFApiResponse
	Data struct {
		FinalExpressRequireCode string                       `json:"finalExpressRequireCode"` // 最终收寄要求CODE
		Details                 []SopCollectShipStandardData `json:"details"`                 // 托寄物收寄标准 list
	} `json:"data"`
}

// 14.产品运费时效推荐接口 IPE_QUERY_PRODUCT_FREIGHT
type IpeQueryProductFreightPost struct {
	ActualWeight     float64       `json:"actualWeight"`               // 实际重量
	Currency         string        `json:"currency"`                   // 声明价值币种，默认取寄件国的默认币种
	CustomerType     int           `json:"customerType"`               // 客户类型 1-个人用户 2-企业用户
	DeclaredValue    int           `json:"declaredValue"`              // 保价声明价值
	DestContIdentity null.Int      `json:"destContIdentity,omitempty"` // 收方件种属性代码 （1-个人 2-公司 默认1）
	DestAddress      null.String   `json:"destAddress,omitempty"`      // 收件地址，省市区+详细地址
	DestCityCode     string        `json:"destCityCode"`               // 目的地城市编码
	DestCountry      string        `json:"destCountry"`                // 目的国二字码，注意：港澳台是HK，MO及TW
	DestPostcode     null.String   `json:"destPostcode,omitempty"`     // 收件邮编
	DestRegionFirst  null.String   `json:"destRegionFirst,omitempty"`  // 一级区划
	DestRegionSecond null.String   `json:"destRegionSecond,omitempty"` // 二级区划
	DestRegionThird  null.String   `json:"destRegionThird,omitempty"`  // 三级区划
	OperateType      int           `json:"operateType"`                // 操作环节类型 1：下单；2：揽收
	ParcelNum        int           `json:"parcelNum"`                  // 包裹数
	SerialNumber     string        `json:"serialNumber"`               // 序列号（建议唯一ID，用于后续定位问题）
	SourceAddress    null.String   `json:"sourceAddress,omitempty"`    // 寄件地址，省市区+详细地址
	SourceCityCode   string        `json:"sourceCityCode"`             // 原寄递城市编码
	SourceCountry    string        `json:"sourceCountry"`              // 原寄国二字码，注意：港澳台是HK，MO及TW
	SourcePostcode   null.String   `json:"sourcePostcode,omitempty"`   // 寄件邮编
	Cargos           []CargoInfo   `json:"cargos,omitempty"`           // 货物集合
	Products         []ProductInfo `json:"products,omitempty"`         // 产品集合
}

type CargoInfo struct {
	FirstClassify  null.String `json:"firstClassify,omitempty"`  // 一级分类
	Name           string      `json:"name"`                     // 品名
	SecondClassify null.String `json:"secondClassify,omitempty"` // 二级分类
	ThirdClassify  null.String `json:"thirdClassify,omitempty"`  // 三级分类
}

type ProductInfo struct {
	InterProductCode null.String `json:"interProductCode"` // 国际产品代码 若有值则优先取该值
}

type IpeQueryProductFreightData struct{}

type ProductListDto struct {
	Calculate   CalculateDto `json:"calculate"`   // 提示级别
	Product     ProductDto   `json:"product"`     // 提示信息
	TimePromise string       `json:"timePromise"` // 提示类型编码
}

type CalculateDto struct {
	Currency string  `json:"currency"` // 提示级别
	Freight  float64 `json:"freight"`  // 提示信息
}

type ProductDto struct {
	CargoTypeCode               string `json:"cargoTypeCode"`               // 快件内容代码
	ExpressTypeCode             string `json:"expressTypeCode"`             // 业务类型代码
	InterProductCode            string `json:"interProductCode"`            // 国际产品代码
	LimitTypeCode               string `json:"limitTypeCode"`               // 时效类型代码
	ProductCode                 string `json:"productCode"`                 // 产品代码L2
	ProductCodeChild            string `json:"productCodeChild"`            // 产品代码L3
	ProductNameChild            string `json:"productNameChild"`            // 产品名称L3
	ProductNameChildExtend      string `json:"productNameChildExtend"`      // 产品名称L4
	ProductNameChildExtendAlias string `json:"productNameChildExtendAlias"` // 产品名称L4 别称
	ProductNo                   string `json:"productNo"`                   // 扩展代码
	ProductSeriesCode           string `json:"productSeriesCode"`           // 主产品代码L1
}

type IssuesDto struct {
	Level    string `json:"level"`    // 提示级别
	Msg      string `json:"msg"`      // 提示信息
	TypeCode string `json:"typeCode"` // 提示类型编码
}

type IpeQueryProductFreightResp struct {
	SFApiResponse
	Data struct {
		Products []ProductListDto `json:"products"` // 产品信息
		Issues   []IssuesDto      `json:"issues"`   // 全局问题配项
	} `json:"data"`
}
