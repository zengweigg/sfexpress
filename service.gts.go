package sfexpress

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/model"
)

type gtsService service

// GtsQueryTrack 路由查询
func (s gtsService) GtsQueryTrack(postData model.GtsQueryTrackPost) (model.GtsQueryTrackResp, error) {
	respData := model.GtsQueryTrackResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetHeaders(map[string]string{
			"lang": "en",
		}).
		SetBody(postData).
		Post("GTS_QUERY_TRACK")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsSendProveOnefile 单票派送证明查询
func (s gtsService) SendProveOnefile(postData model.IomsSendProveOnefilePost) (model.IomsSendProveOnefileResp, error) {
	respData := model.IomsSendProveOnefileResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_SEND_PROVE_ONEFILE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsSendProveCreate 批量派送证明查询
func (s gtsService) SendProveCreate(postData model.IomsSendProveCreatePost) (model.IomsSendProveCreateResp, error) {
	// 步骤1. 调用【查询运单文件接口】，创建查询任务，并生成taskCode返回。
	respData := model.IomsSendProveCreateResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_SEND_PROVE_CREATE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsSendProveCreate 批量派送证明查询
func (s gtsService) SendProveQuery(postData model.IomsSendProveQueryPost) (model.IomsSendProveQueryResp, error) {
	// 步骤2. 调用【步骤1】得到taskCode三分钟后，通过taskCode调用【查询任务结果接口】获取结果。
	respData := model.IomsSendProveQueryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_SEND_PROVE_CREATE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsQueryCharge 计费重和费用项查询
func (s gtsService) QueryCharge(postData model.IomsQueryChargePost) (model.IomsQueryChargeResp, error) {
	respData := model.IomsQueryChargeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_QUERY_CHARGE")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// AdesAgentQueryPod 尾程供应商派送证明查询
func (s gtsService) AdesAgentQueryPod(postData model.AdesAgentQueryPodPost) (model.AdesAgentQueryPodResp, error) {
	respData := model.AdesAgentQueryPodResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ADES_AGENT_QUERY_POD")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsTaxPayment 获取税金信息及支付渠道
func (s gtsService) TaxPayment(postData model.IomsTaxPaymentPost) (model.IomsTaxPaymentResp, error) {
	respData := model.IomsTaxPaymentResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_TAX_PAYMENT")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsTaxPaymentLink 获取税金支付链接
func (s gtsService) TaxPaymentLink(postData model.IomsTaxPaymentLinkPost) (model.IomsTaxPaymentLinkResp, error) {
	respData := model.IomsTaxPaymentLinkResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_TAX_PAYMENT_LINK")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsPaymentLink 获取在线支付短链接
func (s gtsService) PaymentLink(postData model.IomsPaymentLinkPost) (model.IomsPaymentLinkResp, error) {
	respData := model.IomsPaymentLinkResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_PAYMENT_LINK")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// AdesUploadPod 派送证明上传
func (s gtsService) AdesUploadPod(postData model.AdesUploadPodPost) (model.AdesUploadPodResp, error) {
	respData := model.AdesUploadPodResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ADES_UPLOAD_POD")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IomsWaybillBill 操作运单费用明细查询
func (s gtsService) WaybillBill(postData model.IomsWaybillBillPost) (model.IomsWaybillBillResp, error) {
	respData := model.IomsWaybillBillResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IOMS_WAYBILL_BILL")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// CnTaxQuery 中国海关税金信息查询
func (s gtsService) CnTaxQuery(postData model.ICSMCnTaxQueryPost) (model.ICSMCnTaxQueryResp, error) {
	respData := model.ICSMCnTaxQueryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICSM_CN_TAX_QUERY")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// AdesAgentUpdateWaybill 代理商更新运单
func (s gtsService) AdesAgentUpdateWaybill(postData model.AdesAgentUpdateWaybillPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ADES_AGENT_UPDATE_WAYBILL")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// SopCollectShipStandard 国际收寄标准
func (s gtsService) SopCollectShipStandard(postData model.SopCollectShipStandardPost) (model.SopCollectShipStandardResp, error) {
	respData := model.SopCollectShipStandardResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("SOP_COLLECT_SHIP_STANDARD")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// IpeQueryProductFreight 产品运费时效推荐接口
func (s gtsService) IpeQueryProductFreight(postData model.IpeQueryProductFreightPost) (model.IpeQueryProductFreightResp, error) {
	respData := model.IpeQueryProductFreightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IPE_QUERY_PRODUCT_FREIGHT")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
