package sfexpress

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/model"
)

type iecsService service

// CreateOrder 出口电商业务场景接入-创建物流订单
// https://openapi-portal.sf.global/document?firstMenu=API&secondMenu=97
// IECS出口电商业务通常仅有CN、HK始发出口流向
// 注意：创建物流订单之后，不支持修改订单。如果需要修改订单，需先调用"取消电商订单"接口，再重新创建物订单。
func (s iecsService) CreateOrder(postData model.IECSCreateOrderPost) (model.IECSCreateOrderResp, error) {
	respData := model.IECSCreateOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_CREATE_ORDER")
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

// PrintOrder 获取发货面单
func (s iecsService) PrintOrder(postData model.IECSPrintOrderPost) (model.IECSPrintOrderResp, error) {
	respData := model.IECSPrintOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_PRINT_ORDER")
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

// ImportPackage 导入包票关系
func (s iecsService) ImportPackage(postData model.IECSImportPackagePost) (model.IECSImportPackageResp, error) {
	respData := model.IECSImportPackageResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_IMPORT_PACKAGE")
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

// QueryWeightFreight 重量及运费查询
func (s iecsService) QueryWeightFreight(postData model.IECSQueryWeightFreightPost) (model.IECSQueryWeightFreightResp, error) {
	respData := model.IECSQueryWeightFreightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_QUERY_WEIGHT_FREIGHT")
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

// UploadWeight 小包上传重量
func (s iecsService) UploadWeight(postData model.IECSUploadWeightPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_UPLOAD_WEIGHT")
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

// QueryAgentno 服务商单号查询
func (s iecsService) QueryAgentno(postData model.IECSQueryAgentnoPost) (model.IECSQueryAgentnoResp, error) {
	respData := model.IECSQueryAgentnoResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_QUERY_AGENTNO")
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

// CancelOrder 取消电商订单
func (s iecsService) CancelOrder(postData model.IECSCancelOrderPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_CANCEL_ORDER")
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

// OrderIntercept 订单拦截
func (s iecsService) OrderIntercept(postData model.IECSOrderInterceptPost) (model.SFApiResponse, error) {
	respData := model.SFApiResponse{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_ORDER_INTERCEPT")
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

// QueryPod 电商发货/签收证明
func (s iecsService) QueryPod(postData model.IECSQueryPodPost) (model.IECSQueryPodResp, error) {
	respData := model.IECSQueryPodResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("IECS_QUERY_POD")
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
