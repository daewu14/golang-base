package borzo

import (
	"go_base_project/app/repositories/borzo"
	"go_base_project/app/response"
)

type ShowOrderService struct {
	OrderID string
	Repo    borzo.BorzoOrderRepository
}

func (service ShowOrderService) Do() response.ServiceResponse {

	data, err := service.Repo.ShowOrder(service.OrderID)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("Ok", data)

}
