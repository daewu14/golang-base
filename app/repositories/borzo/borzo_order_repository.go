package borzo

import (
	"bytes"
	"encoding/json"
	"errors"
	"go_base_project/app/repositories/borzo/models/order"
	"go_base_project/app/repositories/borzo/models/price"
	"go_base_project/app/repositories/borzo/models/responses"
	"go_base_project/app/services/borzo/base"
	"io"
)

type BorzoOrderRepoInterface interface {
	Order(data order.OrderData) (responses.ResponseDataCreateOrder, error)
	ShowOrder(orderID string) (responses.ResponseDataShowOrder, error)
	Price(data price.DataPrice) (responses.ResponseDataPricing, error)
	CancelOrder()
	ListOrder()
}

type BorzoOrderRepository struct {
}

func (repo BorzoOrderRepository) Order(data order.OrderData) (responses.ResponseDataCreateOrder, error) {

	call, err := base.BorzoApi{}.Post("api/business/1.1/create-order").Bodys(data).Call()

	if err != nil {
		return responses.ResponseDataCreateOrder{}, err
	}

	result, errRa := io.ReadAll(call.Body)

	if errRa != nil {
		return responses.ResponseDataCreateOrder{}, errRa
	}

	var responseDataCreateOrder responses.ResponseDataCreateOrder

	errUm := json.Unmarshal(result, &responseDataCreateOrder)

	if responseDataCreateOrder.IsSuccessful != true {
		errorString := bytes.NewBuffer(result).String()
		return responses.ResponseDataCreateOrder{}, errors.New(errorString)
	}

	if errUm != nil {
		return responses.ResponseDataCreateOrder{}, errUm
	}

	return responseDataCreateOrder, nil
}

func (repo BorzoOrderRepository) Price(data price.DataPrice) (responses.ResponseDataPricing, error) {

	call, err := base.BorzoApi{}.Post("api/business/1.1/calculate-order").Bodys(data).Call()
	if err != nil {
		return responses.ResponseDataPricing{}, err
	}

	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return responses.ResponseDataPricing{}, errRa
	}

	var responseDataPricing responses.ResponseDataPricing
	errUm := json.Unmarshal(result, &responseDataPricing)
	if responseDataPricing.IsSuccessful != true {
		errorString := bytes.NewBuffer(result).String()
		return responses.ResponseDataPricing{}, errors.New(errorString)
	}

	if errUm != nil {
		return responses.ResponseDataPricing{}, errUm
	}

	return responseDataPricing, nil

}

func (repo BorzoOrderRepository) CancelOrder() {
	//TODO implement me
	panic("implement me")
}

func (repo BorzoOrderRepository) ListOrder() {
	//TODO implement me
	panic("implement me")
}

func (repo BorzoOrderRepository) ShowOrder(orderID string) (responses.ResponseDataShowOrder, error) {

	var params = map[string]string{
		"order_id": orderID,
	}

	call, err := base.BorzoApi{}.Get("api/business/1.1/orders").Params(params).Call()

	if err != nil {
		return responses.ResponseDataShowOrder{}, err
	}

	result, errRa := io.ReadAll(call.Body)

	if errRa != nil {
		return responses.ResponseDataShowOrder{}, errRa
	}

	var responseData responses.ResponseDataShowOrder

	errUm := json.Unmarshal(result, &responseData)

	if responseData.IsSuccessful != true {
		errorString := bytes.NewBuffer(result).String()
		return responses.ResponseDataShowOrder{}, errors.New(errorString)
	}

	if errUm != nil {
		return responses.ResponseDataShowOrder{}, errUm
	}

	return responseData, nil
}
