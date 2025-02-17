package sfexpress

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/model"
)

type iuopService service

// CreateOrder 创建物流订单
func (s gtsService) CreateOrder(postData model.IUOPCreateOrderPost) (model.IUOPCreateOrderResp, error) {
	respData := model.IUOPCreateOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_CREATE_ORDER")
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

// QueryOrder 查询物流订单
func (s gtsService) QueryOrder(postData model.IUOPQueryOrderPost) (model.IUOPQueryOrderResp, error) {
	respData := model.IUOPQueryOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_QUERY_ORDER")

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

// PrintOrder 获取发货面单
func (s gtsService) PrintOrder(postData model.IUOPPrintOrderPost) (model.IUOPPrintOrderResp, error) {
	respData := model.IUOPPrintOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_PRINT_ORDER")

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

// CancelOrder 取消物流订单
func (s gtsService) CancelOrder(postData model.IUOPCancelOrderPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_CANCEL_ORDER")

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

// UploadCertify 上传清关资料
func (s gtsService) UploadCertify(postData model.IUOPUploadCertifyPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_UPLOAD_CERTIFY")

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

// ValidityCertify 清关资料是否有效查询
func (s gtsService) ValidityCertify(postData model.IUOPValidityCertifyPost) (model.IUOPValidityCertifyResp, error) {
	respData := model.IUOPValidityCertifyResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_VALIDITY_CERTIFY")

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

// RequireCertify 上传清关资料要求查询
func (s gtsService) RequireCertify(postData model.IUOPRequireCertifyPost) (model.IUOPRequireCertifyResp, error) {
	respData := model.IUOPRequireCertifyResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_REQUIRE_CERTIFY")

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

// QueryFreight 预估运费
func (s gtsService) QueryFreight(postData model.IUOPQueryFreightPost) (model.IUOPQueryFreightResp, error) {
	respData := model.IUOPQueryFreightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_QUERY_FREIGHT")

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

// AsynPrintOrder 异步打印面单
func (s gtsService) AsynPrintOrder(postData model.IUOPAsynPrintOrderPost) (model.IUOPAsynPrintOrderResp, error) {
	respData := model.IUOPAsynPrintOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_ASYN_PRINT_ORDER")

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

// AsynPrintResult 异步打印面单结果查询
func (s gtsService) AsynPrintResult(postData model.IUOPAsynPrintResultPost) (model.IUOPAsynPrintResultResp, error) {
	respData := model.IUOPAsynPrintResultResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_ASYN_PRINT_RESULT")

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

// EstimateFee 预估总费用
func (s gtsService) EstimateFee(postData model.IUOPEstimateFeePost) (model.IUOPEstimateFeeResp, error) {
	respData := model.IUOPEstimateFeeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_ESTIMATE_FEE")

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

// BbdPreMerge BBD预合单
func (s gtsService) BbdPreMerge(postData model.IUOPBbdPreMergePost) (model.IUOPBbdPreMergeResp, error) {
	respData := model.IUOPBbdPreMergeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_BBD_PRE_MERGE")

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

// BbdMergeResult BBD合单结果查询
func (s gtsService) BbdMergeResult(postData model.IUOPBbdMergeResultPost) (model.IUOPBbdMergeResultResp, error) {
	respData := model.IUOPBbdMergeResultResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_BBD_MERGE_RESULT")

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

// CancelBbdMerge 取消BBD预合单
func (s gtsService) CancelBbdMerge(postData model.IUOPCancelBbdMergePost) (model.IUOPCancelBbdMergeResp, error) {
	respData := model.IUOPCancelBbdMergeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_CANCEL_BBD_MERGE")

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

// PreOrder 预创建物流订单
func (s gtsService) PreOrder(postData model.IUOPPreOrderPost) (model.IUOPPreOrderResp, error) {
	respData := model.IUOPPreOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IUOP_PRE_ORDER")

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
