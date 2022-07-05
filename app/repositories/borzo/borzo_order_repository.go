package borzo

import (
	"bytes"
	"encoding/json"
	"errors"
	"go_base_project/app/repositories/borzo/models/order"
	"go_base_project/app/repositories/borzo/models/reponses"
	"go_base_project/app/services/borzo/base"
	"io"
)

type BorzoOrderRepoInterface interface {
	Order(data order.OrderData) (reponses.ResponseDataCreateOrder, error)
	ShowOrder(orderID string) (reponses.ResponseDataShowOrder, error)
	Price()
	CancelOrder()
	ListOrder()
}

type BorzoOrderRepository struct {
}

func (repo BorzoOrderRepository) Order(data order.OrderData) (reponses.ResponseDataCreateOrder, error) {

	call, err := base.BorzoApi{}.Post("api/business/1.1/create-order").Bodys(data).Call()

	if err != nil {
		return reponses.ResponseDataCreateOrder{}, err
	}

	result, errRa := io.ReadAll(call.Body)

	if errRa != nil {
		return reponses.ResponseDataCreateOrder{}, errRa
	}

	var responseDataCreateOrder reponses.ResponseDataCreateOrder

	errUm := json.Unmarshal(result, &responseDataCreateOrder)

	if responseDataCreateOrder.IsSuccessful != true {
		errorString := bytes.NewBuffer(result).String()
		return reponses.ResponseDataCreateOrder{}, errors.New(errorString)
	}

	if errUm != nil {
		return reponses.ResponseDataCreateOrder{}, errUm
	}

	return responseDataCreateOrder, nil
}

func (repo BorzoOrderRepository) Price() {
	//TODO implement me
	panic("implement me")
}

func (repo BorzoOrderRepository) CancelOrder() {
	//TODO implement me
	panic("implement me")
}

func (repo BorzoOrderRepository) ListOrder() {
	//TODO implement me
	panic("implement me")
}

func (repo BorzoOrderRepository) ShowOrder(orderID string) (reponses.ResponseDataShowOrder, error) {

	var params = map[string]string{
		"order_id": orderID,
	}

	call, err := base.BorzoApi{}.Get("api/business/1.1/orders").Params(params).Call()

	if err != nil {
		return reponses.ResponseDataShowOrder{}, err
	}

	result, errRa := io.ReadAll(call.Body)

	if errRa != nil {
		return reponses.ResponseDataShowOrder{}, errRa
	}

	var responseData reponses.ResponseDataShowOrder

	errUm := json.Unmarshal(result, &responseData)

	if responseData.IsSuccessful != true {
		errorString := bytes.NewBuffer(result).String()
		return reponses.ResponseDataShowOrder{}, errors.New(errorString)
	}

	if errUm != nil {
		return reponses.ResponseDataShowOrder{}, errUm
	}

	return responseData, nil
}
