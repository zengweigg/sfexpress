package sfexpress

// 关务

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/sfexpress/model"
)

type icsmService service

// InvoiceInfoUpload 发票信息接收
func (s icsmService) InvoiceInfoUpload(postData model.ICSMInvoiceInfoUploadPost) (model.ICSMInvoiceInfoUploadResp, error) {
	respData := model.ICSMInvoiceInfoUploadResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICSM_INVOICE_INFO_UPLOAD")
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

// InvoiceFileUpload 发票文件接收
func (s icsmService) InvoiceFileUpload(postData model.ICSMInvoiceFileUploadPost) (model.ICSMInvoiceFileUploadResp, error) {
	respData := model.ICSMInvoiceFileUploadResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ICSM_INVOICE_FILE_UPLOAD")
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
