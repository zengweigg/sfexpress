package model

import "gopkg.in/guregu/null.v4"

// 1.创建物流订单 IUOP_CREATE_ORDER
type IUOPCreateOrderPost struct {
	CustomerCode          string                 `json:"customerCode"`          // 客户编码
	Version               null.String            `json:"version"`               // 版本号（默认为空即可）
	OrderOperateType      null.Int               `json:"orderOperateType"`      // 订单操作类型：1-新增 5-改单 默认1
	CustomerOrderNo       string                 `json:"customerOrderNo"`       // 客户订单号 （不能与历史订单号重复）
	CustomerOrderNoTwo    null.String            `json:"customerOrderNoTwo"`    // 客户订单号2
	SfWaybillNo           null.String            `json:"sfWaybillNo"`           // 顺丰运单号
	InterProductCode      string                 `json:"interProductCode"`      // 国际产品映射码，即顺丰物流服务产品代码 参考附录国际产品映射码声明
	ParcelQuantity        int                    `json:"parcelQuantity"`        // 包裹总件数(大于1表示是子母件，最大值300)
	ParcelTotalWeight     null.Float             `json:"parcelTotalWeight"`     // 包裹总重量。精确到小数点后2位，默认为KG
	ParcelWeightUnit      null.String            `json:"parcelWeightUnit"`      // 包裹总重量单位，默认KG
	ParcelTotalLength     null.Float             `json:"parcelTotalLength"`     // 客户订单货物总长。精确到小数点后2位，默认单位CM
	ParcelTotalWidth      null.Float             `json:"parcelTotalWidth"`      // 客户订单货物总宽。精确到小数点后2位，默认单位CM
	ParcelTotalHeight     null.Float             `json:"parcelTotalHeight"`     // 客户订单货物总高。精确到小数点后2位，默认单位CM
	DeclaredValue         float64                `json:"declaredValue"`         // 总商品申报价值，数值应等于ParcelInfo中所有物品的价值之和
	DeclaredCurrency      string                 `json:"declaredCurrency"`      // 申报价值币种 ，参考附录币种声明
	ParcelVolumeUnit      null.String            `json:"parcelVolumeUnit"`      // 客户订单货物长度单位 可选项：CM、INCH 原寄国为CA和US时，长度单位为INCH，其余国家/地区为CM
	PickupType            null.String            `json:"pickupType"`            // 寄件方式：0-服务点自寄或自行联系快递员 1-上门收件
	PickupAppointTime     null.String            `json:"pickupAppointTime"`     // 上门取件预约时间。格式：yyyy-MM-dd HH:mm 若寄件方式选择了上门收件，则预约方式必填
	PickupAppointTimeZone null.String            `json:"pickupAppointTimeZone"` // 上门预约时间时区 详细参考附录字典
	Remark                null.String            `json:"remark"`                // 运单备注
	IsBbd                 null.String            `json:"isBbd"`                 // 是否BBD 0-不是 1-BBD单（需提前开通BBD权限）
	OrderExtendInfo       OrderExtendInfo        `json:"orderExtendInfo"`       // 订单额外信息
	PaymentInfo           PaymentInfo            `json:"paymentInfo"`           // 付款方式信息
	SenderInfo            IuopSenderInfo         `json:"senderInfo"`            // 寄件人信息
	ReceiverInfo          IuopReceiverInfo       `json:"receiverInfo"`          // 收件人信息
	CustomsInfo           IuopCustomsInfo        `json:"customsInfo"`           // 报关信息
	ParcelInfoList        []IuopParcelInfo       `json:"parcelInfoList"`        // 包裹信息
	AddServiceInfoList    []AddServiceInfo       `json:"addServiceInfoList"`    // 增值服务信息
	ChildInfoList         []ChildInfo            `json:"childInfoList"`         // 子单信息
	ExtendAttributeList   []OrderExtendAttribute `json:"extendAttributeList"`   // 扩展属性，具体字段见附录 扩展属性
	EcommerceInfo         OrderEcommerceInfo     `json:"ecommerceInfo"`         // 电子商务信息
	AgentWaybillNo        null.String            `json:"agentWaybillNo"`        // 代理运单号
	LabelModeSize         null.Int               `json:"labelModeSize"`         // 面单打印样式；默认为1
	IsReturnRouteLabel    null.Int               `json:"isReturnRouteLabel"`    // 是否返回路由标签：1：返回路由标签 0：不返回
}

type OrderExtendBase struct {
	IsSignBack          null.String `json:"isSignBack,omitempty"`
	IsSelfPick          null.String `json:"isSelfPick,omitempty"`
	SignBackWaybillNo   null.String `json:"signBackWaybillNo,omitempty"`
	PreMergeSfWaybillNo null.String `json:"preMergeSfWaybillNo,omitempty"`
}

type OrderExtendInfo struct {
	OrderExtendBase
	PoNumber            null.String `json:"poNumber,omitempty"`
	ImporterCompanyName null.String `json:"importerCompanyName,omitempty"`
	ImporterPhoneCode   null.String `json:"importerPhoneCode,omitempty"`
	ImporterPhoneNumber null.String `json:"importerPhoneNumber,omitempty"`
	ImporterAddress     null.String `json:"importerAddress,omitempty"`
	ImportCompTaxNo     null.String `json:"importCompTaxNo,omitempty"`
	ImportCompContacts  null.String `json:"importCompContacts,omitempty"`
	ImportCompEmail     null.String `json:"importCompEmail,omitempty"`
}

type PaymentInfo struct {
	PayMethod       string      `json:"payMethod"`
	PayMonthCard    null.String `json:"payMonthCard,omitempty"`
	TaxPayMethod    string      `json:"taxPayMethod"`
	TaxPayMonthCard null.String `json:"taxPayMonthCard,omitempty"`
}

type IuopSenderInfo struct {
	Company       null.String `json:"company,omitempty"`       // 寄件方公司
	Contact       string      `json:"contact"`                 // 寄件方姓名
	Country       string      `json:"country"`                 // 寄件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	PostCode      string      `json:"postCode"`                // 邮编
	RegionFirst   string      `json:"regionFirst"`             // 寄件方地址一级区划（州/省）
	RegionSecond  string      `json:"regionSecond"`            // 寄件方地址二级区划（城市）
	RegionThird   null.String `json:"regionThird,omitempty"`   // 寄件方地址三级区划（区/县）
	Address       string      `json:"address"`                 // 寄件方详细地址
	Email         null.String `json:"email,omitempty"`         // 寄件方邮箱
	TelAreaCode   null.String `json:"telAreaCode,omitempty"`   // 寄件方固定电话区号
	TelNo         null.String `json:"telNo,omitempty"`         // 寄件方固定电话。寄件方手机号和寄件方固定电话至少填写一个
	PhoneAreaCode null.String `json:"phoneAreaCode,omitempty"` // 寄件方移动电话区号
	PhoneNo       null.String `json:"phoneNo,omitempty"`       // 寄件方移动电话。寄件方手机号和寄件方固定电话至少填写一个
	CargoType     null.Int    `json:"cargoType,omitempty"`     // 寄件类型：1-个人件；2-公司件（CN/TW/KR始发必填）
	CertType      null.String `json:"certType,omitempty"`      // 寄件方证件类型 参考附录证件类型声明
	CertCardNo    null.String `json:"certCardNo,omitempty"`    // 寄件方证件号
	Vat           null.String `json:"vat,omitempty"`           // 寄件方VAT号
	Eori          null.String `json:"eori,omitempty"`          // 寄件方EORI号
	IossNo        null.String `json:"iossNo,omitempty"`        // 寄件方IOSS号
	LvgNo         null.String `json:"lvgNo,omitempty"`         // 寄件方LVG号（马来西亚进口专用，单票500RM以内，过去12个月销售额超过50万RM使用）
}

type IuopReceiverInfo struct {
	Company       null.String `json:"company,omitempty"`       // 收件方公司
	Contact       string      `json:"contact"`                 // 收件方姓名
	Country       string      `json:"country"`                 // 收件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	PostCode      string      `json:"postCode"`                // 邮编
	RegionFirst   string      `json:"regionFirst"`             // 收件方地址一级区划（州/省）
	RegionSecond  string      `json:"regionSecond"`            // 收件方地址二级区划（城市）
	RegionThird   null.String `json:"regionThird,omitempty"`   // 收件方地址三级区划（区/县）
	Address       string      `json:"address"`                 // 收件方详细地址
	RegionFifth   null.String `json:"regionFifth,omitempty"`   // 收件方地址五级区划（道路名，收件方国家为韩国，则必填，长度100）
	RegionSixth   null.String `json:"regionSixth,omitempty"`   // 收件方地址六级区划（建筑编号，收件方国家为韩国，则必填）
	Email         null.String `json:"email,omitempty"`         // 收件方邮箱（若收件国家/地区为GB|IE|NL|MT|DE|FR|ES|IT|BE|AT|HU|PT|DK|CZ|HR|GR|BG|LU|SI|SK|RO|CY|SE|FI|LT|LV|PL|EE|IN|CH|VN 时必填）
	TelAreaCode   null.String `json:"telAreaCode,omitempty"`   // 收件方固定电话区号
	TelNo         null.String `json:"telNo,omitempty"`         // 收件方固定电话。收件方手机号和收件方固定电话至少填写一个
	PhoneAreaCode null.String `json:"phoneAreaCode,omitempty"` // 收件方移动电话区号
	PhoneNo       null.String `json:"phoneNo,omitempty"`       // 收件方移动电话。收件方手机号和收件方固定电话至少填写一个
	CargoType     null.Int    `json:"cargoType,omitempty"`     // 收件类型：1-个人件 2-公司件
	CertType      null.String `json:"certType,omitempty"`      // 收件方证件类型 参考附录证件类型声明 （跨境件收件国家/地区为ID,TW时必填）
	CertCardNo    null.String `json:"certCardNo,omitempty"`    // 收件方证件号 （跨境件收件国家/地区为ID,TW时必填）
	Vat           null.String `json:"vat,omitempty"`           // 收件方VAT号
	Eori          null.String `json:"eori,omitempty"`          // 收件方EORI号
	ContactI18n   null.String `json:"contactI18n,omitempty"`   // 收件方姓名（多语言）
	AddressI18n   null.String `json:"addressI18n,omitempty"`   // 收件方详细地址（多语言）
}

type IuopCustomsInfo struct {
	TradeCondition      null.String `json:"tradeCondition,omitempty"`      // 贸易条件，参考附录字典
	SenderReason        null.Int    `json:"senderReason,omitempty"`        // 寄件原因：1-商业 2-非商业 5-礼物 6-销售 7-样品 8-维修（需联系顺丰特殊配置） 12-其它
	SenderReasonContent null.String `json:"senderReasonContent,omitempty"` // 寄件原因内容
	BusinessRemark      null.String `json:"businessRemark,omitempty"`      // 商业发票备注
	CustomsBatch        null.String `json:"customsBatch,omitempty"`        // 报关批次
	ImportCustomsType   null.String `json:"importCustomsType,omitempty"`   // 进口报关方式 参考附录字典
	ExportCustomsType   null.String `json:"exportCustomsType,omitempty"`   // 出口报关方式 参考附录字典
	HarmonizedCode      null.String `json:"harmonizedCode,omitempty"`      // Harmonized Code美国寄出的可填写
	AesNo               null.String `json:"aesNo,omitempty"`               // AES No.美国寄出可填写
}

type IuopParcelInfo struct {
	Name               string      `json:"name"`                         // 商品名称
	EName              null.String `json:"eName,omitempty"`              // 商品英文名称
	Unit               string      `json:"unit"`                         // 货物单位，如：个，台
	Amount             float64     `json:"amount"`                       // 单件商品申报价值
	Quantity           float64     `json:"quantity"`                     // 数量
	OriginCountry      string      `json:"originCountry"`                // 原产国家，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	HsCode             null.String `json:"hsCode,omitempty"`             // 海关编码
	Brand              null.String `json:"brand,omitempty"`              // 品牌
	StateBarCode       null.String `json:"stateBarCode,omitempty"`       // 国条码
	ProductCustomsNo   null.String `json:"productCustomsNo,omitempty"`   // 海关备案号
	ProductRecordNo    null.String `json:"productRecordNo,omitempty"`    // 货物产品国检备案编号
	GoodsCode          null.String `json:"goodsCode,omitempty"`          // 商品编号
	GoodsUrl           null.String `json:"goodsUrl,omitempty"`           // 商品链接
	Weight             null.Float  `json:"weight,omitempty"`             // 托寄物重量，单件商品SKU重量（不能小于0, 单位KG）。支持小数点后2位
	GoodsDesc          null.String `json:"goodsDesc,omitempty"`          // 商品描述
	Material           null.String `json:"material,omitempty"`           // 材质
	Specifications     null.String `json:"specifications,omitempty"`     // 型号
	ItemNo             null.String `json:"itemNo,omitempty"`             // 商品序号
	QtyOne             null.Float  `json:"qtyOne,omitempty"`             // 法定第一数量
	UnitOne            null.String `json:"unitOne,omitempty"`            // 法定第一计量单位
	QtyTwo             null.Float  `json:"qtyTwo,omitempty"`             // 法定第二数量
	UnitTwo            null.String `json:"unitTwo,omitempty"`            // 法定第二计量单位
	HtsCode            null.String `json:"htsCode,omitempty"`            // 海关编码(美国)
	HtsDesc            null.String `json:"htsDesc,omitempty"`            // 海关编码描述
	Eccn               null.String `json:"eccn,omitempty"`               // 出口管制号码
	ProductNameElement null.String `json:"productNameElement,omitempty"` // 品名申报要素，支持多个要素，要素名使用“[]”标记，多要素使用“;”分隔，填写规范："[{"brand":"test "},{"specifications":"test2"},{"mode":"L3"}]"
	SellerVatNumber    null.String `json:"sellerVatNumber,omitempty"`    // 税号
}

type MaterialInfo struct {
	MaterialCode string `json:"materialCode,omitempty"` // 物料编码
	MaterialNum  int    `json:"materialNum,omitempty"`  // 物料数量
}

type AddServiceBaseInfo struct {
	ServiceCode null.String `json:"serviceCode,omitempty"` // 增值服务编码（参考附录增值服务声明）
	Value       null.String `json:"value,omitempty"`       // 增值服务扩展属性
	ValueOne    null.String `json:"valueOne,omitempty"`    // 增值服务扩展属性1
}

type AddServiceInfo struct {
	AddServiceBaseInfo
	ValueTwo         null.String    `json:"valueTwo,omitempty"`         // 增值服务扩展属性2
	ValueThree       null.String    `json:"valueThree,omitempty"`       // 增值服务扩展属性3
	ValueFour        null.String    `json:"valueFour,omitempty"`        // 增值服务扩展属性4
	ValueFive        null.String    `json:"valueFive,omitempty"`        // 增值服务扩展属性5
	MaterialInfoList []MaterialInfo `json:"materialInfoList,omitempty"` // 包装服务物料信息
}

type ChildBaseInfo struct {
	SfWaybillNo string `json:"sfWaybillNo,omitempty"` // 子单编号
}

type ChildInfo struct {
	SfWaybillNo null.String `json:"sfWaybillNo,omitempty"` // 子单编号
	Weight      null.Float  `json:"weight,omitempty"`      // 托寄物重量
	Length      null.Float  `json:"length,omitempty"`      // 长，单位厘米
	Width       null.Float  `json:"width,omitempty"`       // 宽，单位厘米
	Height      null.Float  `json:"height,omitempty"`      // 高
	BoxNo       null.String `json:"boxNo,omitempty"`       // 箱号；箱号在子单集合中不能重复；如果使用箱号，请把包含母单对应的箱号全部传入
}

type OrderExtendAttribute struct {
	AttributeName  null.String `json:"attributeName,omitempty"`  // 属性名 参考附录字典
	AttributeValue null.String `json:"attributeValue,omitempty"` // 属性值
}

type OrderEcommerceInfo struct {
	PayerCertType            null.String `json:"payerCertType,omitempty"`            // 订购人证件类型
	PayerCertNo              null.String `json:"payerCertNo,omitempty"`              // 订购人证件号码
	PayerName                null.String `json:"payerName,omitempty"`                // 订购人姓名
	PayerTel                 null.String `json:"payerTel,omitempty"`                 // 订购人联络电话
	BuyerEmail               null.String `json:"buyerEmail,omitempty"`               // 买家邮件地址
	BuyerNickname            null.String `json:"buyerNickname,omitempty"`            // 买家昵称
	BuyerFreight             null.Float  `json:"buyerFreight,omitempty"`             // 买家邮费（电商收取消费者的邮费）
	BuyerPremium             null.Float  `json:"buyerPremium,omitempty"`             // 买家保费（电商收取消费者的保价费用）
	Tax                      null.String `json:"tax,omitempty"`                      // 税款，企业预先代扣的税款金额
	PreferentialAmount       null.Float  `json:"preferentialAmount,omitempty"`       // 优惠金额（整个订单的，电商给消费者的优惠金额）
	ActPay                   null.Float  `json:"actPay,omitempty"`                   // 成交实际付款金额（消费者向电商支付的金额，商品价格+运杂费+代扣税款-优惠金额）
	Currency                 null.String `json:"currency,omitempty"`                 // 结算币种
	PayToolType              null.String `json:"payToolType,omitempty"`              // 支付工具、方式，如支付宝，财付通
	ApplicationNo            null.String `json:"applicationNo,omitempty"`            // 申请单编号
	PayNo                    null.String `json:"payNo,omitempty"`                    // 支付交易号
	PlatformDomainName       null.String `json:"platformDomainName,omitempty"`       // 电商平台域名
	PayTime                  null.Time   `json:"payTime,omitempty"`                  // 支付交易时间
	PromotionType            null.String `json:"promotionType,omitempty"`            // 促销类型
	SellerEmail              null.String `json:"sellerEmail,omitempty"`              // 卖家邮件地址
	SellerNickname           null.String `json:"sellerNickname,omitempty"`           // 卖家昵称
	SellerAlipayNo           null.String `json:"sellerAlipayNo,omitempty"`           // 卖家支付宝账号
	PayEnterpriseCustomsCode null.String `json:"payEnterpriseCustomsCode,omitempty"` // 支付企业的海关注册登记编号
	PayEnterpriseCheckCode   null.String `json:"payEnterpriseCheckCode,omitempty"`   // 支付企业的国检注册登记编号
	PayEnterpriseName        null.String `json:"payEnterpriseName,omitempty"`        // 支付企业名称
	StoreCode                null.String `json:"storeCode,omitempty"`                // 店铺代码
	ShopWebCode              null.String `json:"shopWebCode,omitempty"`              // 购物网站代码
	EnterpriseName           null.String `json:"enterpriseName,omitempty"`           // 电商企业的海关注册登记名称
	EnterpriseCustomsCode    null.String `json:"enterpriseCustomsCode,omitempty"`    // 电商企业的海关注册登记编号
	EnterpriseCheckCode      null.String `json:"enterpriseCheckCode,omitempty"`      // 电商企业的国检注册登记编号
	EnterpriseCrossCode      null.String `json:"enterpriseCrossCode,omitempty"`      // 电商企业的跨境平台注册登记编号
	PlatformName             null.String `json:"platformName,omitempty"`             // 电商平台的海关注册登记名称
	PlatformCustomsCode      null.String `json:"platformCustomsCode,omitempty"`      // 电商平台的海关注册登记编号
	PlatformCheckCode        null.String `json:"platformCheckCode,omitempty"`        // 电商平台的国检注册登记编号
	PartZoneCode             null.String `json:"partZoneCode,omitempty"`             // 分区码
	PortCode                 null.String `json:"portCode,omitempty"`                 // 口岸
	ImportCustomsType        null.String `json:"importCustomsType,omitempty"`        // 进口报关方式
}

type IuopRespChildInfo struct {
	SfWaybillNo string  `json:"sfWaybillNo,omitempty"` // 子单编号
	Weight      float64 `json:"weight,omitempty"`      // 托寄物重量
	Length      float64 `json:"length,omitempty"`      // 长，单位厘米
	Width       float64 `json:"width,omitempty"`       // 宽，单位厘米
	Height      float64 `json:"height,omitempty"`      // 高
	BoxNo       string  `json:"boxNo,omitempty"`       // 箱号；箱号在子单集合中不能重复；如果使用箱号，请把包含母单对应的箱号全部传入
}

type IuopRespRouteLabelInfo struct {
	DestRouteLabel   string `json:"destRouteLabel"`   // 路由标签
	TwoDimensionCode string `json:"twoDimensionCode"` // 面单二维码信息
}

type IUOPCreateOrderResp struct {
	SFApiResponse
	Data struct {
		CustomerOrderNo    string                   `json:"customerOrderNo"`    // 订单号
		SfWaybillNo        string                   `json:"sfWaybillNo"`        // 运单号
		AgentWaybillNo     string                   `json:"agentWaybillNo"`     // 代理/服务商运单号
		ChildWaybillNoList []string                 `json:"childWaybillNoList"` // 子运单号
		IdentityUploadUrl  string                   `json:"identityUploadUrl"`  // 身份证上传地址
		InvoiceUrl         string                   `json:"invoiceUrl"`         // 发票打印地址
		LabelUrl           string                   `json:"labelUrl"`           // 运单打印地址
		RouteLabelInfoList []IuopRespRouteLabelInfo `json:"routeLabelInfoList"` // 路由标签信息
		ChildInfoList      []IuopRespChildInfo      `json:"childInfoList"`      // 子件信息
	} `json:"data"`
}

// 2.查询物流订单 IUOP_QUERY_ORDER
type IUOPQueryOrderPost struct {
	CustomerCode    string      `json:"customerCode"`              // 客户编码
	Version         null.String `json:"version,omitempty"`         // 版本号（默认为空即可）
	SfWaybillNo     null.String `json:"sfWaybillNo,omitempty"`     // 顺丰运单号(母单号)，客户订单号customerOrderNo、 顺丰运单号sfWayBillNo至少填一个，如果两个都填写，校验是否为同一订单
	CustomerOrderNo null.String `json:"customerOrderNo,omitempty"` // 客户订单号，客户订单号customerOrderNo、 顺丰运单号sfWayBillNo至少填一个，如果两个都填写，校验是否为同一订单
}

type ParcelInfo2 struct {
	Name             string  `json:"name"`             // 商品名称
	EName            string  `json:"eName"`            // 商品英文名称
	Unit             string  `json:"unit"`             // 货物单位
	Amount           float64 `json:"amount"`           // 单价
	Quantity         float64 `json:"quantity"`         // 数量
	OriginCountry    string  `json:"originCountry"`    // 原产国家
	HsCode           string  `json:"hsCode"`           // 海关编码
	Brand            string  `json:"brand"`            // 品牌
	StateBarCode     string  `json:"stateBarCode"`     // 国条码
	ProductCustomsNo string  `json:"productCustomsNo"` // 海关备案号
	ProductRecordNo  string  `json:"productRecordNo"`  // 货物产品国检备案编号
	GoodsCode        string  `json:"goodsCode"`        // 商品编号
	GoodsUrl         string  `json:"goodsUrl"`         // 商品链接
	Weight           float64 `json:"weight"`           // 托寄物重量
	GoodsDesc        string  `json:"goodsDesc"`        // 商品描述
	Material         string  `json:"material"`         // 材质
	Specifications   string  `json:"specifications"`   // 型号
}

type IUOPQueryOrderResp struct {
	SFApiResponse
	Data struct {
		CustomerOrderNo       string             `json:"customerOrderNo"`       // 客户订单号
		CustomerOrderNoTwo    string             `json:"customerOrderNoTwo"`    // 客户订单号2
		SfWaybillNo           string             `json:"sfWaybillNo"`           // 顺丰运单号
		InterProductCode      string             `json:"interProductCode"`      // 国际产品映射码
		ParcelQuantity        int                `json:"parcelQuantity"`        // 包裹总件数(大于1表示是子母件，最大值300)
		ParcelTotalWeight     float64            `json:"parcelTotalWeight"`     // 包裹总重量。精确到小数点后2位
		ParcelWeightUnit      string             `json:"parcelWeightUnit"`      // 包裹总重量单位
		ParcelTotalLength     float64            `json:"parcelTotalLength"`     // 客户订单货物总长。精确到小数点后2位
		ParcelTotalWidth      float64            `json:"parcelTotalWidth"`      // 客户订单货物总宽。精确到小数点后2位
		ParcelTotalHeight     float64            `json:"parcelTotalHeight"`     // 客户订单货物总高。精确到小数点后2位
		DeclaredValue         string             `json:"declaredValue"`         // 申明价值
		DeclaredCurrency      string             `json:"declaredCurrency"`      // 包裹总计声明价值币种
		ParcelVolumeUnit      string             `json:"parcelVolumeUnit"`      // 客户订单货物长度单位 1.可选项：CM、INCH 2.默认单位为CM 3.原寄国为CA和US，长度单位为英寸
		PickupType            string             `json:"pickupType"`            // 寄件方式 0-服务点自寄或自行联系快递员； 1-上门收件；
		PickupAppointTime     string             `json:"pickupAppointTime"`     // 上门取件预约时间。格式：yyyy-MM-dd HH:mm 若寄件方式选择了上门收件，则预约方式必填
		PickupAppointTimeZone string             `json:"pickupAppointTimeZone"` // 上门预约时间时区 详细参考附录字典
		Remark                string             `json:"remark"`                // 运单备注
		OrderExtendInfo       OrderExtendBase    `json:"orderExtendInfo"`       // 订单额外信息
		PaymentInfo           PaymentInfo        `json:"paymentInfo"`           // 付款方式信息
		SenderInfo            SenderInfo2        `json:"senderInfo"`            // 寄件人信息
		ReceiverInfo          ReceiverInfo2      `json:"receiverInfo"`          // 收件人信息
		CustomsInfo           IuopCustomsInfo    `json:"customsInfo"`           // 报关信息
		ParcelInfoList        ParcelInfo2        `json:"parcelInfoList"`        // 包裹信息
		AddServiceInfoList    AddServiceBaseInfo `json:"addServiceInfoList"`    // 增值服务信息
		ChildInfoList         []ChildBaseInfo    `json:"childInfoList"`         // 子单信息
		AgentWaybillNo        string             `json:"agentWaybillNo"`        // 代理运单号
	} `json:"data"`
}

// 3.获取发货面单 IUOP_PRINT_ORDER
type IUOPPrintOrderPost struct {
	CustomerCode          string               `json:"customerCode,omitempty"`          // 客户编码
	RequestId             string               `json:"requestId,omitempty"`             // 请求id（可使用UUID，确保请求唯一性字段，只要确保每个请求的requestId唯一即可）
	Version               null.String          `json:"version,omitempty"`               // 版本号（默认为空即可）
	LabelModeSize         null.Int             `json:"labelModeSize,omitempty"`         // 面单打印样式
	PrintType             string               `json:"printType,omitempty"`             // 打印类型。1-运单 2-发票
	PrintWaybillNoDtoList []PrintWaybillNoDto2 `json:"printWaybillNoDtoList,omitempty"` // 支持批量打印 母单相同的子单（含母单） 打印参数对象
	FileFormat            string               `json:"fileFormat,omitempty"`            // 打印文件格式：PDF或者ZPL；默认为PDF
	Resolution            string               `json:"resolution,omitempty"`            // ZPL文件分辨率，支持203dpi和300dpi，仅在fileFormat为ZPL时有效
}

type PrintWaybillNoDto2 struct {
	SfWaybillNo              string   `json:"sfWaybillNo"`                        // 顺丰运单号
	IsPrintSubParent         null.Int `json:"isPrintSubParent,omitempty"`         // 是否打印全部的子母单 0：否 1：是 默认值：0
	SinglePrintWaybillNoList []string `json:"singlePrintWaybillNoList,omitempty"` // 单独打印的运单号集合(含母单，子单)
}

type IUOPPrintOrderData struct {
	Url string `json:"url"`
}

type IUOPPrintOrderResp struct {
	SFApiResponse
	Data IUOPPrintOrderData `json:"data"`
}

// 4.取消物流订单 IUOP_CANCEL_ORDER
type IUOPCancelOrderPost struct {
	CustomerCode    string      `json:"customerCode"`              // 客户编码
	SfWaybillNo     null.String `json:"sfWaybillNo,omitempty"`     // 顺丰运单号(母单号),顺丰运单号与客户订单号二者填其一
	CustomerOrderNo null.String `json:"customerOrderNo,omitempty"` // 客户订单号,顺丰运单号与客户订单号二者填其一
}

// 5.上传清关资料 IUOP_UPLOAD_CERTIFY
type IUOPUploadCertifyPost struct {
	CustomerCode    string      `json:"customerCode"`              // 客户编码
	Version         null.String `json:"version,omitempty"`         // 版本号（默认为空即可）
	SfWaybillNo     null.String `json:"sfWaybillNo,omitempty"`     // 顺丰运单号(母单号),顺丰运单号与客户订单号二者填其一
	CustomerOrderNo null.String `json:"customerOrderNo,omitempty"` // 客户订单号,顺丰运单号与客户订单号二者填其一
	CertType        string      `json:"certType"`                  // 参考附录 “证件类型声明” 以及 “报关资料类型”
	FileFormat      null.String `json:"fileFormat,omitempty"`      // 文件格式（JPG、PNG、PDF）
	FileContent     null.String `json:"fileContent,omitempty"`     // 文件内容（Base64编码） 若证件类型是身份证，此字段建议为身份证正面
	FileContentTwo  null.String `json:"fileContentTwo,omitempty"`  // 文件内容（Base64编码） 若证件类型是身份证，此字段建议为身份证反面
	CertName        null.String `json:"certName,omitempty"`        // 证件姓名
	CertCardNo      null.String `json:"certCardNo,omitempty"`      // 证件号码
	CertsFlag       null.String `json:"certsFlag,omitempty"`       // 证件归属：0-收方，1-寄方
}

// 6.清关资料是否有效查询 IUOP_VALIDITY_CERTIFY
type IUOPValidityCertifyPost struct {
	Version  null.String `json:"version,omitempty"`  // 版本号（默认为空即可）
	UserName string      `json:"userName"`           // 客户姓名
	CardNo   null.String `json:"cardNo,omitempty"`   // 证件号
	TelPhone null.String `json:"telPhone,omitempty"` // 客户电话
}

type IUOPValidityCertifyResp struct {
	SFApiResponse
	Data bool `json:"data"`
}

// 7.上传清关资料要求查询 IUOP_REQUIRE_CERTIFY
type IUOPRequireCertifyPost struct {
	CustomerCode    string      `json:"customerCode"`              // 客户编码
	Version         null.String `json:"version,omitempty"`         // 版本号（默认为空即可）
	RequestId       string      `json:"requestId"`                 // 请求id（可使用UUID，确保请求唯一性字段，只要确保每个请求的requestId唯一即可）
	SfWaybillNo     null.String `json:"sfWaybillNo,omitempty"`     // 顺丰运单号，客户订单号和顺丰运单号不能同时为空
	CustomerOrderNo null.String `json:"customerOrderNo,omitempty"` // 客户订单号，客户订单号和顺丰运单号不能同时为空
}

type CusCertsDetailDto struct {
	CertNameId      string `json:"certNameId"`      // 单证名称ID
	CertName        string `json:"certName"`        // 单证名称(需要根据语言)
	CertDescription string `json:"certDescription"` // 单证描述(需要根据语言)
	TelUpload       string `json:"telUpload"`       // 是否上传号码（1：上传，0：无需上传）
	ImageUpload     string `json:"imageUpload"`     // 是否上传图片（0：无需上传，1：需上传，2：已上传需审核， 3：已上传无需审核）
	CertAmount      int    `json:"certAmount"`      // 上传单证图片的最小数量(上传图片时必填)
	CertMaxAmount   int    `json:"certMaxAmount"`   // 上传单证图片的最大数量(上传图片时必填)
}

type CusCertGroupInfoDto struct {
	CertState             string              `json:"certState"`             // 单证提示说明（需要根据语言）
	NeedNums              int                 `json:"needNums"`              // 所需单证名称数量 (needNums=-1，表示仅提示；needNums=0，表示选填；needNums=2，表示这几个里面必须上传2项；)
	CusCertsDetailDtolist []CusCertsDetailDto `json:"cusCertsDetailDtolist"` // 单证名称
}

type CusCertsInfoDto struct {
	IeFlag                  string                `json:"ieFlag"`                  // 进出口标识
	CertType                string                `json:"certType"`                // 单证类型（报关资料(I) 报关证件(C)）
	CertsFlag               string                `json:"certsFlag"`               // 单证归属, 0-收方,1-寄方
	CusCertGroupInfoDtoList []CusCertGroupInfoDto `json:"cusCertGroupInfoDtoList"` // 需上传证件详细规则
}

type UploadStatisticsCertDto struct {
	CertNameId        string   `json:"certNameId"`        // 单证名称ID
	CertName          string   `json:"certName"`          // 单证名称(需要根据语言)
	PostFlag          string   `json:"postFlag"`          // 单证归属, 0-收方,1-寄方
	CertAmount        int      `json:"certAmount"`        // 单证数量
	CertNo            string   `json:"certNo"`            // 上传的证件号
	OtherCertNameList []string `json:"otherCertNameList"` // 其他报关资料上传的文件名集合
}

type IUOPRequireCertifyResp struct {
	SFApiResponse
	Data struct {
		CusCertsInfoList      []CusCertsInfoDto         `json:"cusCertsInfoList"`      // 需上传证件的要求
		UploadedCertsInfoList []UploadStatisticsCertDto `json:"uploadedCertsInfoList"` // 已上传证件
		Tips                  string                    `json:"tips"`                  // 温馨提示
	} `json:"data"`
}

// 8.预估运费 IUOP_QUERY_FREIGHT
type IUOPQueryFreightPost struct {
	CustomerCode     string      `json:"customerCode"`           // 客户编码
	Version          null.String `json:"version,omitempty"`      // 版本号（默认为空即可）
	SenderCountry    string      `json:"senderCountry"`          // 寄件国家/地区的二字码 采用ISO 3166 标准 Alpha-2 code 如：CN,US,JP,KR,SG,MY
	ReceiverCountry  string      `json:"receiverCountry"`        // 收件国家/地区的二字码 采用ISO 3166 标准 Alpha-2 code 如：CN,US,JP,KR,SG,MY
	SenderPostCode   string      `json:"senderPostCode"`         // 寄方邮编
	ReceiverPostCode string      `json:"receiverPostCode"`       // 到方邮编
	InterProductCode string      `json:"interProductCode"`       // 产品统一映射码
	PayMonthCard     null.String `json:"payMonthCard,omitempty"` // 月结账号
	ConsignedTm      null.String `json:"consignedTm,omitempty"`  // 寄件时间,YYYY-MM-DD HH:mm:ss（默认当前时间）
	Weight           string      `json:"weight"`                 // 实际重量（KG）
	Currency         null.String `json:"currency,omitempty"`     // 币种，传入则按该币种返回费用
}

type IUOPQueryFreightResp struct {
	SFApiResponse
	Data struct {
		Currency       string `json:"currency"`       // 币种
		TotalFreight   string `json:"totalFreight"`   // 总费用
		Weight         string `json:"weight"`         // 重量
		DiscountedRate string `json:"discountedRate"` // 折扣运费
	} `json:"data"`
}

// 9.异步打印面单 IUOP_ASYN_PRINT_ORDER
type IUOPAsynPrintOrderPost struct {
	CustomerCode          string               `json:"customerCode"`          // 客户编码
	RequestId             string               `json:"requestId"`             // 请求id（可使用UUID，确保请求唯一性字段，只要确保每个请求的requestId唯一即可）
	Version               null.String          `json:"version"`               // 版本号（默认为空即可）
	LabelModeSize         null.String          `json:"labelModeSize"`         // 面单打印样式
	PrintType             string               `json:"printType"`             // 打印类型。1-运单 2-发票
	PrintWaybillNoDtoList []PrintWaybillNoDto2 `json:"printWaybillNoDtoList"` // 支持批量打印 母单相同的子单（含母单） 打印参数对象
	FileFormat            string               `json:"fileFormat"`            // 打印文件格式：PDF或者ZPL；默认为PDF
	Resolution            string               `json:"resolution"`            // ZPL文件分辨率，支持203dpi和300dpi，仅在fileFormat为ZPL时有效
}

type IUOPAsynPrintOrderResp struct {
	SFApiResponse
	Data struct {
		PrintRequestId string `json:"printRequestId"`
	} `json:"data"`
}

// 10.异步打印面单结果查询 IUOP_ASYN_PRINT_RESULT
type IUOPAsynPrintResultPost struct {
	CustomerCode   string      `json:"customerCode"`   // 客户编码
	RequestId      string      `json:"requestId"`      // 请求id（可使用UUID，确保请求唯一性字段，只要确保每个请求的requestId唯一即可）
	Version        null.String `json:"version"`        // 版本号（默认为空即可）
	PrintRequestId string      `json:"printRequestId"` // 打印请求id
}

type IUOPAsynPrintResultResp struct {
	SFApiResponse
	Data IUOPPrintOrderData `json:"data"`
}

// 11.预估总费用 IUOP_ESTIMATE_FEE
type IUOPEstimateFeePost struct {
	CustomerCode       string            `json:"customerCode"`       // 客户编码
	Version            null.String       `json:"version"`            // 版本号（默认为空即可）
	InterProductCode   string            `json:"interProductCode"`   // 国际产品映射码
	ParcelQuantity     null.Int          `json:"parcelQuantity"`     // 包裹总件数(大于1表示是子母件，最大值300)
	ParcelTotalWeight  null.Float        `json:"parcelTotalWeight"`  // 包裹总重量。精确到小数点后2位
	ParcelTotalLength  null.Float        `json:"parcelTotalLength"`  // 客户订单货物总长。精确到小数点后2位
	ParcelTotalWidth   null.Float        `json:"parcelTotalWidth"`   // 客户订单货物总宽。精确到小数点后2位
	ParcelTotalHeight  null.Float        `json:"parcelTotalHeight"`  // 客户订单货物总高。精确到小数点后2位
	DeclaredValue      null.Float        `json:"declaredValue"`      // 总商品申报价值，单件商品申报价值*数量获取
	DeclaredCurrency   null.String       `json:"declaredCurrency"`   // 申报价值币种
	PaymentInfo        PaymentInfo       `json:"paymentInfo"`        // 付款方式信息
	SenderInfo         SenderInfo3       `json:"senderInfo"`         // 寄件人信息
	ReceiverInfo       ReceiverInfo3     `json:"receiverInfo"`       // 收件人信息
	CustomsInfo        CustomsInfo2      `json:"customsInfo"`        // 报关信息
	AddServiceInfoList []AddServiceInfo3 `json:"addServiceInfoList"` // 增值服务信息
}

type ReceiverInfo3 struct {
	Country      string      `json:"country"`      // 收件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	PostCode     string      `json:"postCode"`     // 邮编
	RegionFirst  null.String `json:"regionFirst"`  // 收件方地址一级区划（州/省）
	RegionSecond null.String `json:"regionSecond"` // 收件方地址二级区划（城市）
	RegionThird  null.String `json:"regionThird"`  // 收件方地址三级区划（区/县）
	Address      null.String `json:"address"`      // 收件方详细地址
	CargoType    null.Int    `json:"cargoType"`    // 收件方货物类型：1-个人件；2-公司件
}

type CustomsInfo2 struct {
	ImportCustomsType null.String `json:"importCustomsType"` // 进口报关方式 参考附录字典
	ExportCustomsType null.String `json:"exportCustomsType"` // 出口报关方式 参考附录字典
}

type AddServiceInfo3 struct {
	ServiceCode string      `json:"serviceCode"` // 增值服务代码
	Value       string      `json:"value"`       // 保价申报价值
	ValueOne    null.String `json:"valueOne"`    // 增值服务扩展属性1
}

type IUOPEstimateFeeData struct {
	ServiceName    string  `json:"serviceName"`    // 费用类型
	EstimateFee    float64 `json:"estimateFee"`    // 预估费用金额
	EstimatePayFee float64 `json:"estimatePayFee"` // 折后费用金额
}

type IUOPEstimateFeeResp struct {
	SFApiResponse
	Data struct {
		Currency         string                `json:"currency"`         // 币种
		TotalFee         float64               `json:"totalFee"`         // 总费用
		Weight           float64               `json:"weight"`           // 重量
		InterProductCode string                `json:"interProductCode"` // 国际产品映射代码
		FeeInfoList      []IUOPEstimateFeeData `json:"feeInfoList"`      // 费用信息列表
	} `json:"data"`
}

// 12.BBD预合单 IUOP_BBD_PRE_MERGE
type IUOPBbdPreMergePost struct {
	CustomerCode string      `json:"customerCode"` // 客户编码
	Version      null.String `json:"version"`      // 版本号（默认为空即可）
	SfWayBillNos []string    `json:"sfWayBillNos"` // 待合运单号
}

type IUOPBbdPreMergeResp struct {
	SFApiResponse
	Data struct {
		PreMergeSfWaybillNo string `json:"preMergeSfWaybillNo"` // 预合单号
		PrintUrl            string `json:"printUrl"`            // 运单打印地址
	} `json:"data"`
}

// 13.BBD合单结果查询 IUOP_BBD_MERGE_RESULT
type IUOPBbdMergeResultPost struct {
	CustomerCode     string `json:"customerCode"`     // 客户编码
	SfWaybillNo      string `json:"sfWaybillNo"`      // 顺丰运单号（母单或子单）
	MergeSfWaybillNo string `json:"mergeSfWaybillNo"` // 合单单号（母单或子单）
}

type BBDDtoInfo struct {
	MergeSfWaybillNo string `json:"mergeSfWaybillNo"` // 合单号（大单母单号或子单号）
	SfWaybillNos     string `json:"sfWaybillNos"`     // 顺丰运单号（可包含多个，格式：运单号,运单号,运单号,...）
}

type IUOPBbdMergeResultData struct {
	MergeSfWaybillNos string       `json:"mergeSfWaybillNos"` // 合单单号（可包含多个，格式：母单号,子单号,子单号...）
	MergeRelationList []BBDDtoInfo `json:"mergeRelationList"` // 合单关系
}

type IUOPBbdMergeResultResp struct {
	SFApiResponse
	Data []IUOPBbdMergeResultData `json:"data"`
}

// 14.取消BBD预合单 IUOP_CANCEL_BBD_MERGE
type IUOPCancelBbdMergePost struct {
	CustomerCode string   `json:"customerCode"` // 客户编码
	SfWayBillNos []string `json:"sfWayBillNos"` // 预合单号列表
}

type IUOPCancelBbdMergeResp struct {
	SFApiResponse
	Data bool `json:"data"`
}

// 15.预创建物流订单 IUOP_PRE_ORDER
type IUOPPreOrderPost struct {
	CustomerCode          string                 `json:"customerCode"`          // 客户编码
	Version               null.String            `json:"version"`               // 版本号（默认为空即可）
	OrderOperateType      null.Int               `json:"orderOperateType"`      // 订单操作类型
	CustomerOrderNo       string                 `json:"customerOrderNo"`       // 客户订单号 （不能与历史订单号重复）
	CustomerOrderNoTwo    null.String            `json:"customerOrderNoTwo"`    // 客户订单号2
	SfWaybillNo           null.String            `json:"sfWaybillNo"`           // 顺丰运单号
	InterProductCode      string                 `json:"interProductCode"`      // 国际产品映射码，即顺丰物流服务产品代码
	ParcelQuantity        int                    `json:"parcelQuantity"`        // 包裹总件数(大于1表示是子母件，最大值300)
	ParcelTotalWeight     null.Float             `json:"parcelTotalWeight"`     // 包裹总重量。精确到小数点后2位，默认为KG
	ParcelWeightUnit      null.String            `json:"parcelWeightUnit"`      // 包裹总重量单位，默认KG
	ParcelTotalLength     null.Float             `json:"parcelTotalLength"`     // 客户订单货物总长。精确到小数点后2位，默认单位CM
	ParcelTotalWidth      null.Float             `json:"parcelTotalWidth"`      // 客户订单货物总宽。精确到小数点后2位，默认单位CM
	ParcelTotalHeight     null.Float             `json:"parcelTotalHeight"`     // 客户订单货物总高。精确到小数点后2位，默认单位CM
	DeclaredValue         float64                `json:"declaredValue"`         // 总商品申报价值，数值应等于ParcelInfo中所有物品的价值之和
	DeclaredCurrency      string                 `json:"declaredCurrency"`      // 申报价值币种 ，参考附录币种声明
	ParcelVolumeUnit      null.String            `json:"parcelVolumeUnit"`      // 客户订单货物长度单位
	PickupType            null.String            `json:"pickupType"`            // 寄件方式
	PickupAppointTime     null.String            `json:"pickupAppointTime"`     // 上门取件预约时间。格式：yyyy-MM-dd HH:mm
	PickupAppointTimeZone null.String            `json:"pickupAppointTimeZone"` // 上门预约时间时区 详细参考附录字典
	Remark                null.String            `json:"remark"`                // 运单备注
	IsBbd                 null.String            `json:"isBbd"`                 // 是否BBD 0-不是 1-BBD单（需提前开通BBD权限）
	OrderExtendInfo       OrderExtendInfo        `json:"orderExtendInfo"`       // 订单额外信息
	PaymentInfo           PaymentInfo            `json:"paymentInfo"`           // 付款方式信息
	SenderInfo            SenderInfo4            `json:"senderInfo"`            // 寄件人信息
	ReceiverInfo          ReceiverInfo4          `json:"receiverInfo"`          // 收件人信息
	CustomsInfo           CustomsInfo4           `json:"customsInfo"`           // 报关信息
	ParcelInfoList        []ParcelInfo4          `json:"parcelInfoList"`        // 包裹信息
	AddServiceInfoList    []AddServiceBaseInfo   `json:"addServiceInfoList"`    // 增值服务信息
	ChildInfoList         []ChildInfo            `json:"childInfoList"`         // 子单信息
	ExtendAttributeList   []OrderExtendAttribute `json:"extendAttributeList"`   // 扩展属性，具体字段见附录 扩展属性
	EcommerceInfo         OrderEcommerceInfo     `json:"ecommerceInfo"`         // 电子商务信息
	AgentWaybillNo        null.String            `json:"agentWaybillNo"`        // 代理运单号
	LabelModeSize         null.Int               `json:"labelModeSize"`         // 面单打印样式；默认为1
	IsReturnRouteLabel    null.Int               `json:"isReturnRouteLabel"`    // 是否返回路由标签
}

type SenderInfo4 struct {
	Company       null.String `json:"company"`       // 寄件方公司
	Contact       string      `json:"contact"`       // 寄件方姓名
	Country       string      `json:"country"`       // 寄件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	PostCode      string      `json:"postCode"`      // 邮编
	RegionFirst   string      `json:"regionFirst"`   // 寄件方地址一级区划（州/省）
	RegionSecond  string      `json:"regionSecond"`  // 寄件方地址二级区划（城市）
	RegionThird   null.String `json:"regionThird"`   // 寄件方地址三级区划（区/县）
	Address       string      `json:"address"`       // 寄件方详细地址
	Email         null.String `json:"email"`         // 寄件方邮箱
	TelAreaCode   null.String `json:"telAreaCode"`   // 寄件方固定电话区号
	TelNo         null.String `json:"telNo"`         // 寄件方固定电话。寄件方手机号和寄件方固定电话至少填写一个
	PhoneAreaCode null.String `json:"phoneAreaCode"` // 寄件方移动电话区号
	PhoneNo       null.String `json:"phoneNo"`       // 寄件方移动电话。寄件方手机号和寄件方固定电话至少填写一个
	CargoType     null.Int    `json:"cargoType"`     // 寄件类型：1-个人件；2-公司件（CN/TW/KR始发必填）
	CertType      null.String `json:"certType"`      // 寄件方证件类型 参考附录证件类型声明
	CertCardNo    null.String `json:"certCardNo"`    // 寄件方证件号
	Vat           null.String `json:"vat"`           // 寄件方VAT号
	Eori          null.String `json:"eori"`          // 寄件方EORI号
	IossNo        null.String `json:"iossNo"`        // 寄件方IOSS号
	LvgNo         null.String `json:"lvgNo"`         // 寄件方LVG号（马来西亚进口专用，单票500RM以内，过去12个月销售额超过50万RM使用）
}

type ReceiverInfo4 struct {
	Company       null.String `json:"company"`       // 收件方公司
	Contact       string      `json:"contact"`       // 收件方姓名
	Country       string      `json:"country"`       // 收件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	PostCode      string      `json:"postCode"`      // 邮编
	RegionFirst   string      `json:"regionFirst"`   // 收件方地址一级区划（州/省）
	RegionSecond  string      `json:"regionSecond"`  // 收件方地址二级区划（城市）
	RegionThird   null.String `json:"regionThird"`   // 收件方地址三级区划（区/县）
	Address       string      `json:"address"`       // 收件方详细地址
	RegionFifth   null.String `json:"regionFifth"`   // 收件方地址五级区划（道路名，收件方国家为韩国，则必填，长度100）
	RegionSixth   null.String `json:"regionSixth"`   // 收件方地址六级区划（建筑编号，收件方国家为韩国，则必填）
	Email         null.String `json:"email"`         // 收件方邮箱
	TelAreaCode   null.String `json:"telAreaCode"`   // 收件方固定电话区号
	TelNo         null.String `json:"telNo"`         // 收件方固定电话。收件方手机号和收件方固定电话至少填写一个
	PhoneAreaCode null.String `json:"phoneAreaCode"` // 收件方移动电话区号
	PhoneNo       null.String `json:"phoneNo"`       // 收件方移动电话。收件方手机号和收件方固定电话至少填写一个
	CargoType     null.Int    `json:"cargoType"`     // 收件类型：1-个人件 2-公司件
	CertType      null.String `json:"certType"`      // 收件方证件类型 参考附录证件类型声明 （跨境件收件国家/地区为ID,TW时必填）
	CertCardNo    null.String `json:"certCardNo"`    // 收件方证件号 （跨境件收件国家/地区为ID,TW时必填）
	Vat           null.String `json:"vat"`           // 收件方VAT号
	Eori          null.String `json:"eori"`          // 收件方EORI号
	ContactI18n   null.String `json:"contactI18n"`   // 收件方姓名（多语言）
	AddressI18n   null.String `json:"addressI18n"`   // 收件方详细地址（多语言）
}

type CustomsInfo4 struct {
	TradeCondition      null.String `json:"tradeCondition"`      // 贸易条件，参考附录字典
	SenderReason        null.Int    `json:"senderReason"`        // 寄件原因：1-商业 2-非商业
	SenderReasonContent null.String `json:"senderReasonContent"` // 寄件原因内容
	BusinessRemark      null.String `json:"businessRemark"`      // 商业发票备注
	CustomsBatch        null.String `json:"customsBatch"`        // 报关批次
	ImportCustomsType   null.String `json:"importCustomsType"`   // 进口报关方式 参考附录字典
	ExportCustomsType   null.String `json:"exportCustomsType"`   // 出口报关方式 参考附录字典
	HarmonizedCode      null.String `json:"harmonizedCode"`      // Harmonized Code美国寄出的可填写
	AesNo               null.String `json:"aesNo"`               // AES No.美国寄出可填写
}

type ParcelInfo4 struct {
	Name               string      `json:"name"`               // 商品名称
	EName              null.String `json:"eName"`              // 商品英文名称
	Unit               string      `json:"unit"`               // 货物单位，如：个，台
	Amount             float64     `json:"amount"`             // 单件商品申报价值
	Quantity           float64     `json:"quantity"`           // 数量
	OriginCountry      string      `json:"originCountry"`      // 原产国家，如：CN,US,JP,KR,SG,MY 更多参考附录字典
	HsCode             null.String `json:"hsCode"`             // 海关编码
	Brand              null.String `json:"brand"`              // 品牌
	StateBarCode       null.String `json:"stateBarCode"`       // 国条码
	ProductCustomsNo   null.String `json:"productCustomsNo"`   // 海关备案号
	ProductRecordNo    null.String `json:"productRecordNo"`    // 货物产品国检备案编号
	GoodsCode          null.String `json:"goodsCode"`          // 商品编号
	GoodsUrl           null.String `json:"goodsUrl"`           // 商品链接
	Weight             null.Float  `json:"weight"`             // 托寄物重量，单件商品SKU重量（不能小于0, 单位KG）。支持小数点后2位
	GoodsDesc          null.String `json:"goodsDesc"`          // 商品描述
	Material           null.String `json:"material"`           // 材质
	Specifications     null.String `json:"specifications"`     // 型号
	ItemNo             null.String `json:"itemNo"`             // 商品序号
	QtyOne             null.Float  `json:"qtyOne"`             // 法定第一数量
	UnitOne            null.String `json:"unitOne"`            // 法定第一计量单位
	QtyTwo             null.Float  `json:"qtyTwo"`             // 法定第二数量
	UnitTwo            null.String `json:"unitTwo"`            // 法定第二计量单位
	HtsCode            null.String `json:"htsCode"`            // 海关编码(美国)
	HtsDesc            null.String `json:"htsDesc"`            // 海关编码描述
	Eccn               null.String `json:"eccn"`               // 出口管制号码
	ProductNameElement null.String `json:"productNameElement"` // 品名申报要素，支持多个要素，要素名使用“[]”标记，多要素使用“;”分隔
	SellerVatNumber    null.String `json:"sellerVatNumber"`    // 税号
}
