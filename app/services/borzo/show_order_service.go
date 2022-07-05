package borzo

import (
	"fmt"
	"go_base_project/app/repositories/borzo"
	"go_base_project/app/response"
)

type ShowOrderService struct {
	OrderID string
	Repo    borzo.BorzoOrderRepoInterface
}

func (service ShowOrderService) Do() response.ServiceResponse {

	data, err := service.Repo.ShowOrder(service.OrderID)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	fmt.Println(data)
	return response.Service().Success("Ok", data)

}
