package borzo

import (
	"go_base_project/app/repositories/borzo"
	"go_base_project/app/repositories/borzo/models/order"
	"go_base_project/app/response"
)

type CreateOrderService struct {
	OrderData order.OrderData
	Repo      borzo.BorzoOrderRepository
}

func (service CreateOrderService) Do() response.ServiceResponse {
	data, err := service.Repo.Order(service.OrderData)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", data)
}
