package sfexpress

import (
	"sfexpress/model"
)

type iecsService service

// CreateOrder 创建电商订单
// https://openapi-portal.sf.global/document?firstMenu=API&secondMenu=97
// IECS出口电商业务通常仅有CN、HK始发出口流向
// 注意：创建物流订单之后，不支持修改订单。如果需要修改订单，需先调用"取消电商订单"接口，再重新创建物订单。
func (s iecsService) CreateOrder(order model.SFIECSOrder) (model.SFIECSOrderResponse, error) {
	respData := model.SFIECSOrderResponse{}

	return respData, nil
}
