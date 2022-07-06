package borzo

import (
	"go_base_project/app/repositories/borzo"
	"go_base_project/app/response"
)

type CancelOrderService struct {
	OrderID string
	Repo    borzo.BorzoOrderRepository
}

func (service CancelOrderService) Do() response.ServiceResponse {

	result, err := service.Repo.CancelOrder(service.OrderID)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", result)

}
