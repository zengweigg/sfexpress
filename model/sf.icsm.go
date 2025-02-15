package model

import (
	"gopkg.in/guregu/null.v4"
)

// 1.发票信息接收 ICSM_INVOICE_INFO_UPLOAD
type ICSMInvoiceInfoUploadPost struct {
	DataSource  string         `json:"dataSource"`  // 1. 数据来源，如BYD
	Status      string         `json:"status"`      // 2. 上传动作，N-新增，R-更新，C-删除
	Bno         string         `json:"bno"`         // 3. SF运单号
	GUID        null.String    `json:"GUID"`        // 4. 接口每次交互的流水号
	InvNo       string         `json:"invNo"`       // 5. 发票号
	InvDate     string         `json:"invDate"`     // 6. 发票日期，格式：YYYYMMDD
	OrderNo     null.String    `json:"orderNo"`     // 7. 采购订单号
	IncoTerm    string         `json:"incoTerm"`    // 8. 贸易条款，根据实际情况填写， 按照 UNECE Recommendation 17 标准
	PaymentTerm string         `json:"paymentTerm"` // 9. 付款条件
	Loading     null.String    `json:"loading"`     // 10. 启运港
	Dicharge    null.String    `json:"dicharge"`    // 11. 目的地港
	Mark        null.String    `json:"mark"`        // 12. 唛头（收货人相关信息、品牌、航班、公司名称等相关信息）
	Category    null.String    `json:"category"`    // 13. 分类
	Shipper     Shipper        `json:"shipper"`     // 14. 寄方对象
	Consignee   Consignee      `json:"consignee"`   // 15. 收方对象
	Goods       []GoodsDetails `json:"goods"`       // 16. 商品对象集合
}

type GoodsDetails struct {
	Description     string      `json:"description"`     // 英文商品描述
	DescriptionThai string      `json:"descriptionThai"` // 泰文商品描述
	Brand           string      `json:"brand"`           // 品牌
	PartNo          string      `json:"partNo"`          // 型号/工厂料号
	Qty             int         `json:"qty"`             // 申报数量
	QtyUnit         string      `json:"qtyUnit"`         // 申报数量单位
	UPrice          float64     `json:"uPrice"`          // 单价
	Amount          float64     `json:"amount"`          // 总价
	Currency        string      `json:"currency"`        // 币种，按照ISO 4217 (RFCUC)标准
	Weight          float64     `json:"weight"`          // 净重
	WeightUnit      string      `json:"weightUnit"`      // 净重单位，英文三字码
	GrossWeight     float64     `json:"grossWeight"`     // 毛重
	GrossWeightUnit string      `json:"grossWeightUnit"` // 毛重单位
	Origin          string      `json:"origin"`          // 原产国，取ISO 3166 (RFICC) 二字码
	HsCode          null.String `json:"hsCode"`          // 商品编码
	Privilege       null.String `json:"privilege"`       // 贸易类型
}

type Shipper struct {
	Name    string      `json:"name"`    // 1. 发货人英文名称
	Address string      `json:"address"` // 2. 发货人地址
	TaxId   null.String `json:"taxId"`   // 3. 出口方税号，出口申报时此字段必填
}

type Consignee struct {
	Name    string      `json:"name"`    // 1. 进口方英文名称
	Address string      `json:"address"` // 2. 进口方地址
	TaxId   null.String `json:"taxId"`   // 3. 进口方税号，进口申报时此字段必填
}

type ICSMInvoiceInfoUploadData struct {
	GUID      string `json:"GUID"`      // 接口请求流水号
	Bno       string `json:"bno"`       // 运单号
	InvNo     string `json:"invNo"`     // 发票号
	Timestamp string `json:"timestamp"` // 时间戳，YYYYMMDDThhmmss
}

type ICSMInvoiceInfoUploadResp struct {
	SFApiResponse
	Data []ICSMInvoiceInfoUploadData `json:"data"`
}

// 2.发票文件接收 ICSM_INVOICE_FILE_UPLOAD
type ICSMInvoiceFileUploadPost struct {
	Bno             string      `json:"bno"`             // SF运单号
	GUID            null.String `json:"GUID"`            // 请求流水号
	InvNo           string      `json:"invNo"`           // 【进口计划】-表头“清关发票号”,新增字段，必填，【进口台账】表头需要同步显示该字段信息
	DocType         string      `json:"docType"`         // 0 = Invoice/Packing
	PayloadContent  string      `json:"payloadContent"`  // 以ZIP格式将文件压缩后Base64编码
	PayloadFileName string      `json:"payloadFileName"` // 以清关发票号命名，类型是other
	Status          string      `json:"status"`          // N = New, R = Revise, C = Cancel
}

type ICSMInvoiceFileUploadResp struct {
	SFApiResponse
	Data struct {
		GUID      string `json:"GUID"`      // 接口请求流水号
		Bno       string `json:"bno"`       // 运单号
		InvNo     string `json:"invNo"`     // 发票号
		Timestamp string `json:"timestamp"` // 时间戳，YYYYMMDDThhmmss
	} `json:"data"`
}
