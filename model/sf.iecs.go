package model

import (
	"gopkg.in/guregu/null.v4"
)

// 1.出口电商业务场景接入-创建物流订单 IECS_CREATE_ORDER
type IECSCreateOrderPost struct {
	ApiUsername        string      `json:"apiUsername"`        // 1. 客户名称（小包用户名，可登录IUOP管理中心获取）
	CustomerOrderNo    string      `json:"customerOrderNo"`    // 2. 客户订单号，要求唯一值，不能重复
	PlatformOrderId    string      `json:"platformOrderId"`    // 3. 平台订单号，如果是第三方ERP平台，可采用平台的订单号，如果没有或不是平台，可同客户订单号
	PlatformCode       string      `json:"platformCode"`       // 4. 平台编码，电商平台简称必须为英文，非电商平台直接用公司英文名
	ErpCode            string      `json:"erpCode"`            // 5. ERP平台code，如果是第三方ERP，请求联系技术提供，如果是自建系统，默认：0000
	PlatformMerchantId null.String `json:"platformMerchantId"` // 6. 用户在电商平台的ID
	SfWaybillNo        null.String `json:"sfWaybillNo"`        // 7. 顺丰运单号
	InterProductCode   string      `json:"interProductCode"`   // 8. 国际产品映射码，即顺丰物流服务产品代码
	ParcelQuantity     int         `json:"parcelQuantity"`     // 9. 包裹总件数(固定为1)
	ParcelTotalWeight  float64     `json:"parcelTotalWeight"`  // 10. 包裹总重量。精确到小数点后3位，默认为KG
	ParcelTotalLength  null.Float  `json:"parcelTotalLength"`  // 11. 客户订单货物总长。精确到小数点后2位，默认单位CM
	ParcelTotalWidth   null.Float  `json:"parcelTotalWidth"`   // 12. 客户订单货物总宽。精确到小数点后2位，默认单位CM
	ParcelTotalHeight  null.Float  `json:"parcelTotalHeight"`  // 13. 客户订单货物总高。精确到小数点后2位，默认单位CM
	DeclaredValue      float64     `json:"declaredValue"`      // 14. 总商品申报价值，单件商品申报价值*数量获取
	DeclaredCurrency   string      `json:"declaredCurrency"`   // 15. 申报价值币种
	OrderInsurance     null.Float  `json:"orderInsurance"`     // 16. 保价服务费（只能为正整数，且单票不能超过5000人民币）
	TaxNo              null.String `json:"taxNo"`              // 17. 税号
	Abn                null.String `json:"abn"`                // 18. 1.在澳洲有注册公司的企业，基于其ABN号走标准流程注册的号码，11位纯数字。 2.欧盟国家流向，请将IOSS填充于此字段。格式为：2位字母开头+数字。3.马来西亚进口LVG销售税号填写
	GstExemptionCode   null.String `json:"gstExemptionCode"`   // 19. 澳洲商业编号
	IsBat              int         `json:"isBat"`              // 20. 是否带电 0：不带电 ;1 带电
	IsLiquid           null.Int    `json:"isLiquid"`           // 21. 是否液体 0非液体 1是液体
	CreateOrderType    string      `json:"createOrderType"`    // 22. 1：小包专递
	//OrderExtendInfo    null.OrderExtendInfo `json:"orderExtendInfo"`    // 23. 订单额外信息
	SenderInfo         SenderInfo   `json:"senderInfo"`         // 24. 寄件人信息
	ReceiverInfo       ReceiverInfo `json:"receiverInfo"`       // 25. 收件人信息
	CustomsInfo        CustomsInfo  `json:"customsInfo"`        // 26. 报关信息
	ParcelInfoList     []ParcelInfo `json:"parcelInfoList"`     // 27. 包裹信息
	AgentWaybillNo     null.String  `json:"agentWaybillNo"`     // 28. 代理运单号
	InputDirectionCode null.String  `json:"inputDirectionCode"` // 29. 合同物流部分产品指定流向代码
}

type SenderInfoBase struct {
	Country      string      `json:"country"`      // 3. 寄件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典 (必填)
	PostCode     string      `json:"postCode"`     // 4. 邮编 (必填)
	RegionFirst  null.String `json:"regionFirst"`  // 5. 寄件方地址一级区划（州/省） (必填)
	RegionSecond null.String `json:"regionSecond"` // 6. 寄件方地址二级区划（城市） (必填)
	RegionThird  null.String `json:"regionThird"`  // 7. 寄件方地址三级区划（区/县） (非必填)
	Address      null.String `json:"address"`      // 8. 寄件方详细地址 (必填)
}

type SenderInfo struct {
	SenderInfoBase
	Company string      `json:"company"` // 1. 寄件方公司 (必填)
	Contact string      `json:"contact"` // 2. 寄件方姓名 (必填)
	TelNo   string      `json:"telNo"`   // 9. 寄件方固定电话。寄件方手机号和寄件方固定电话至少填写一个 (必填)
	PhoneNo string      `json:"phoneNo"` // 10. 寄件方移动电话。寄件方手机号和寄件方固定电话至少填写一个 (必填)
	Email   null.String `json:"email"`   // 11. 寄件方邮箱 (非必填)
}

type SenderInfo2 struct {
	SenderInfo
	TelAreaCode   string `json:"telAreaCode"`   // 寄件方固定电话区号
	PhoneAreaCode string `json:"phoneAreaCode"` // 寄件方移动电话区号
	CargoType     int    `json:"cargoType"`     // 寄件方货物类型：1-个人件；2-公司件
	CertType      string `json:"certType"`      // 寄件方证件类型 详细参考附录字典
	CertCardNo    string `json:"certCardNo"`    // 寄件方证件号
	Vat           string `json:"vat"`           // 寄件方VAT号
	Eori          string `json:"eori"`          // 寄件方EORI号
	IossNo        string `json:"iossNo"`        // 寄件方IOSS号
}

type SenderInfo3 struct {
	SenderInfoBase
	CargoType null.Int `json:"cargoType"` // 寄件方货物类型：1-个人件；2-公司件
}

type ReceiverInfo struct {
	Company       string      `json:"company"`       // 1. 收件方公司 (必填)
	Contact       string      `json:"contact"`       // 2. 收件方姓名 (必填)
	Country       string      `json:"country"`       // 3. 收件方国家/地区，如：CN,US,JP,KR,SG,MY 更多参考附录字典 (必填)
	PostCode      string      `json:"postCode"`      // 4. 邮编 (必填)
	RegionFirst   string      `json:"regionFirst"`   // 5. 收件方地址一级区划（州/省） (必填)
	RegionSecond  string      `json:"regionSecond"`  // 6. 收件方地址二级区划（城市） (必填)
	RegionThird   null.String `json:"regionThird"`   // 7. 收件方地址三级区划（区/县） (非必填)
	Address       string      `json:"address"`       // 8. 收件方详细地址 (必填)
	Email         string      `json:"email"`         // 9. 收件方邮箱 (必填)
	TelNo         string      `json:"telNo"`         // 10. 收件方固定电话。收件方手机号和收件方固定电话至少填写一个 (必填)
	PhoneAreaCode null.String `json:"phoneAreaCode"` // 11. 收件方移动电话区号 (非必填)
	PhoneNo       string      `json:"phoneNo"`       // 12. 收件方移动电话。收件方手机号和收件方固定电话至少填写一个 (必填)
	DestDoorplate null.String `json:"destDoorplate"` // 13. 到方门牌号，德国必填 (非必填)
}

type ReceiverInfo2 struct {
	ReceiverInfo
	RegionFifth string `json:"regionFifth"` // 收件方地址五级区划（道路名，收件方国家为韩国，则必填，长度100）
	RegionSixth string `json:"regionSixth"` // 收件方地址六级区划（建筑编号，收件方国家为韩国，则必填）
	TelAreaCode string `json:"telAreaCode"` // 收件方固定电话区号
	CargoType   int    `json:"cargoType"`   // 收件方货物类型：1-个人件；2-公司件
	CertType    string `json:"certType"`    // 收件方证件类型 详细参考附录字典
	CertCardNo  string `json:"certCardNo"`  // 收件方证件号
	Vat         string `json:"vat"`         // 收件方VAT号
	Eori        string `json:"eori"`        // 收件方EORI号
}

// 报关信息
type CustomsInfo struct {
	PassportId null.String `json:"passportId"` // 1. 韩国个人清关代码，如P561140689980，只接受英文和数字 专线小包韩国流向为必填 (非必填)
	Category   null.String `json:"category"`   // 2. 订单所属品类，用于海关清关 (非必填)
}

// 海关申报信息

type ParcelInfo struct {
	Name               string      `json:"name"`               // 1. 商品名称（默认使用，可以是各个国家的语言名称） (必填)
	CName              string      `json:"cName"`              // 2. 商品中文名称 (必填)
	Unit               string      `json:"unit"`               // 3. 货物单位，如：个，台 (必填)
	Amount             float64     `json:"amount"`             // 4. 单件商品申报价值 (必填)
	Currency           string      `json:"currency"`           // 5. 单价币种 (必填)
	Quantity           float64     `json:"quantity"`           // 6. 数量 (必填)
	OriginCountry      null.String `json:"originCountry"`      // 7. 原产国家 (非必填)
	HsCode             null.String `json:"hsCode"`             // 8. 海关编码 (非必填)
	GoodsUrl           null.String `json:"goodsUrl"`           // 9. 商品链接 (非必填)
	Weight             float64     `json:"weight"`             // 10. 托寄物重量，单件商品SKU重量（不能小于0, 单位KG）。支持小数点后3位 (必填)
	GoodsDesc          null.String `json:"goodsDesc"`          // 11. 商品描述/拣货信息 (非必填)
	SkuID              null.String `json:"skuID"`              // 12. 商品编码SKU (非必填)
	SkuType            null.String `json:"skuType"`            // 13. 商品类型：0:未知。1:普货。2:敏货。3:带电货。 (非必填)
	ProductNameElement null.String `json:"productNameElement"` // 14. 品名申报要素，支持多个要素，要素名使用“[]”标记，多要素使用“,”分隔，填写规范："[{"brand":"test"},{"specifications":"test2"},{"mode":"L3"}]" (非必填)
}

type IECSCreateOrderResp struct {
	SFApiResponse
	Data struct {
		CustomerOrderNo string `json:"customerOrderNo"` // 1. 订单号 (必填)
		SfWaybillNo     string `json:"sfWaybillNo"`     // 2. 运单号 (必填)
		AgentWaybillNo  string `json:"agentWaybillNo"`  // 3. 代理/服务商运单号 (必填)
		DirectionCode   string `json:"directionCode"`   // 4. 流向代码 (必填)
		AgentCode       string `json:"agentCode"`       // 5. 代理编号 (必填)
	} `json:"data"`
}

// 2.获取发货面单 IECS_PRINT_ORDER
type IECSPrintOrderPost struct {
	Version               null.String         `json:"version"`               // 1. 版本号（默认为空即可）
	PrintWaybillNoDtoList []PrintWaybillNoDto `json:"printWaybillNoDtoList"` // 2. 打印参数对象
	ApiUsername           string              `json:"apiUsername"`           // 3. 客户名称（customerCode对应的用户名）
	PrintPicking          null.Bool           `json:"printPicking"`          // 4. 是否打印拣货单，默认：false
	OnePdf                null.Bool           `json:"onePdf"`                // 5. 是否合成一个PDF，默认：true
}

type PrintWaybillNoDto struct {
	SfWaybillNo string `json:"sfWaybillNo"` // 顺丰运单号
}

type IECSPrintOrderResp struct {
	SFApiResponse
	Data struct {
		Url             string `json:"url"`             // 1. 打印的URL
		SfWaybillNo     string `json:"sfWaybillNo"`     // 2. 顺丰单号
		CustomerOrderNo string `json:"customerOrderNo"` // 3. 用户订单号
	} `json:"data"`
}

// 3.导入包票关系 IECS_IMPORT_PACKAGE
type IECSImportPackagePost struct {
	ApiUsername       string      `json:"apiUsername"`       // 1. 客户名称（customerCode对应的用户名）
	PackageCode       null.String `json:"packageCode"`       // 2. 顺丰包号（由业务人员提供）
	PackageCode2      null.String `json:"packageCode2"`      // 3. 海外代理包号（由业务人员提供）
	PackageCode3      null.String `json:"packageCode3"`      // 4. 客户包号（使用该字段请联系业务人员）
	WaybillList       []ReqDto    `json:"waybillList"`       // 5. 运单集合
	ReturnPackageCode null.String `json:"returnPackageCode"` // 6. 是否返回包号，默认值false
}

type ReqDto struct {
	SfWaybillNo    null.String `json:"sfWaybillNo"`    // 1. 顺丰单号
	AgentWaybillNo null.String `json:"agentWaybillNo"` // 2. 海外代理单号
	Weight         null.String `json:"weight"`         // 3. 重量（kg），仅支持小数后3位，超过3位小数，系统自动四舍五入
}

type RespDtoInfo struct {
	SfWaybillNo    string `json:"sfWaybillNo"`    // 1. 顺丰单号
	AgentWaybillNo string `json:"agentWaybillNo"` // 2. 海外代理单号
	Weight         string `json:"weight"`         // 3. 重量（kg），仅支持小数后3位，超过3位小数，系统自动四舍五入
	FailReason     string `json:"failReason"`     // 4. 失败原因（空时表示成功）
}

type IECSImportPackageResp struct {
	SFApiResponse
	Data struct {
		ApiUsername  string        `json:"apiUsername"`  // 1. 客户名称（customerCode对应的用户名）
		PackageCode  string        `json:"packageCode"`  // 2. 顺丰包号
		PackageCode2 string        `json:"packageCode2"` // 3. 海外代理包号
		PackageCode3 string        `json:"packageCode3"` // 4. 客户包号
		WaybillList  []RespDtoInfo `json:"waybillList"`  // 5. 运单集合
	} `json:"data"`
}

// 4.重量及运费查询 IECS_QUERY_WEIGHT_FREIGHT
type IECSQueryWeightFreightPost struct {
	WeightAndFreightList []WeightFreightList `json:"weightAndFreightList"` // 1. 运单信息集合
}

type WeightFreightList struct {
	SfWaybillNo    string `json:"sfWaybillNo"`    // 1. 顺丰运单号
	AgentWaybillNo string `json:"agentWaybillNo"` // 2. 代理运单号
}

type IECSQueryWeightFreightData struct {
	AgentWaybillNo   string  `json:"agentWaybillNo"`   // 1. 代理运单单号
	Message          string  `json:"message"`          // 2. 提示信息
	SfWaybillNo      string  `json:"sfWaybillNo"`      // 3. 顺丰运单号
	Freight          float64 `json:"freight"`          // 4. 运费（单位：RMB）
	Tariff           float64 `json:"tariff"`           // 5. 关税费 （单位：RMB）
	Clearance        float64 `json:"clearance"`        // 6. 清关费 （单位：RMB）
	ClearanceService float64 `json:"clearanceService"` // 7. 清关服务费 （单位：RMB）
	Weight           float64 `json:"weight"`           // 8. 计费重量 （单位：KG）
	RealWeight       float64 `json:"realWeight"`       // 9. 真实重量 （单位：KG）
	WeightTime       string  `json:"weightTime"`       // 10. 称重时间
	DebitTime        string  `json:"debitTime"`        // 11. 计费时间
	AgentCode        string  `json:"agentCode"`        // 12. 渠道代码
	CargoHeight      float64 `json:"cargoHeight"`      // 13. 货物高度cm
	CargoLength      float64 `json:"cargoLength"`      // 14. 货物长度cm
	CargoWidth       float64 `json:"cargoWidth"`       // 15. 货物宽度cm
	DestCountryCode  string  `json:"destCountryCode"`  // 16. 目的地国家
	DestPostCode     string  `json:"destPostCode"`     // 17. 目的地邮编
	MonthCardNo      string  `json:"monthCardNo"`      // 18. 月结卡号
	OrderTime        string  `json:"orderTime"`        // 19. 下单时间,格式:yyyy-MM-dd HH:mm:ss
	PortCode         string  `json:"portCode"`         // 20. 目的地口岸
	ProductName      string  `json:"productName"`      // 21. 产品名称
	ProductType      string  `json:"productType"`      // 22. 产品类型
}

type IECSQueryWeightFreightResp struct {
	SFApiResponse
	Data []IECSQueryWeightFreightData `json:"data"`
}

// 5.小包上传重量 IECS_UPLOAD_WEIGHT
type IECSUploadWeightPost struct {
	WaybillList []IECSUploadWeightPackageInfo `json:"waybillList"` // 1. 运单信息集
}

type IECSUploadWeightPackageInfo struct {
	SfWaybillNo string      `json:"sfWaybillNo"` // 1. 顺丰运单号
	Weight      string      `json:"weight"`      // 2. 包裹重量(kg)（限制0~50）
	Length      null.Float  `json:"length"`      // 3. 长度，单位cm
	Width       null.Float  `json:"width"`       // 4. 宽度，单位cm
	Height      null.Float  `json:"height"`      // 5. 高度，单位cm
	Number      null.String `json:"number"`      // 6. 序号
	ImgUrlList  []string    `json:"imgUrlList"`  // 7. 图片url集合
}

// 6.服务商单号查询 IECS_QUERY_AGENTNO
type IECSQueryAgentnoPost struct {
	Service        string            `json:"service"`        // 1. 请使用默认值：WaybillService
	WaybillList    []AgentnoPostInfo `json:"waybillList"`    // 2. 运单信息集合
	AgentLabelCtrl string            `json:"agentLabelCtrl"` // 3. 是否需要返回代理面单
}

type AgentnoPostInfo struct {
	SfWaybillNo    string      `json:"sfWaybillNo"`    // 1. 顺丰运单号
	OrderNo        string      `json:"orderNo"`        // 2. 客户订单号（下单时使用的customerOrderNo）
	AgentLabelCtrl null.String `json:"agentLabelCtrl"` // 3. 是否需要返回代理面单
}

type AgentnoWrongData struct {
	SfWaybillNo     string `json:"sfWaybillNo"`     // 顺丰运单号
	OrderNo         string `json:"orderNo"`         // 客户订单号（下单时使用的customerOrderNo`)
	Description     string `json:"description"`     // 错误信息描述
	StatusAgentCode string `json:"statusagentcode"` // 错误信息编码
}

type AgentnoSuccessData struct {
	SfWaybillNo    string `json:"sfWaybillNo"`    // 顺丰运单号
	OrderNo        string `json:"orderNo"`        // 客户订单号（下单时使用的customerOrderNo）
	AgentWaybillNo string `json:"agentWaybillNo"` // 代理单号，部分平邮无代理单号
	Agentcode      string `json:"agentcode"`      // 代理编号
	AgentLabelUrl  string `json:"agentLabelUrl"`  // 代理面单URL
}

type IECSQueryAgentnoResp struct {
	SFApiResponse
	Service     string               `json:"service"`     // 默认值：WaybillService
	SuccessData []AgentnoSuccessData `json:"successData"` // 查询成功列表
	WrongData   []AgentnoWrongData   `json:"wrongData"`   // 查询失败列表
}

// 7.取消电商订单 IECS_CANCEL_ORDER
type IECSCancelOrderPost struct {
	SfWaybillNo     string `json:"sfWaybillNo"`     // 顺丰运单号
	CustomerOrderNo string `json:"customerOrderNo"` // 客户订单号
	CancelReason    string `json:"cancelReason"`    // 取消原因
}

// 8.订单拦截 IECS_ORDER_INTERCEPT
type IECSOrderInterceptPost struct {
	InterceptType   string      `json:"interceptType"`   // 拦截类型：1-拦截、2-取消拦截
	Remark          null.String `json:"remark"`          // 拦截原因/取消拦截原因
	CustomerOrderNo string      `json:"customerOrderNo"` // 客户订单号
	SfWaybillNo     string      `json:"sfWaybillNo"`     // 顺丰单号
}

// 9.电商发货/签收证明 IECS_QUERY_POD
type IECSQueryPodPost struct {
	SfWaybillNo string `json:"sfWaybillNo"` // 顺丰单号
	SignType    string `json:"signType"`    // 0-发货证明,1-签收证明
}

type IECSQueryPodResp struct {
	SFApiResponse
	SmartSignatureMessage string `json:"smartSignatureMessage"` // 文件状态描述
	SmartSignatureStatus  string `json:"smartSignatureStatus"`  // 文件状态：1生成成功，2文件生成中，3顺丰内部审核失败，请联系顺丰客服 （对应ECP申请失败）4其他错误，请联系顺丰客户
	SmartSignatureType    string `json:"smartSignatureType"`    // 文件类型：0-发货证明,1-签收证明
	SmartSignatureUrl     string `json:"smartSignatureUrl"`     // 文件链接：已盖章的文件链接
	UserName              string `json:"userName"`              // 用户名
}
