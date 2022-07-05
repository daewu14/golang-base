package borzo

import (
	"go_base_project/app/repositories/borzo"
	"go_base_project/app/repositories/borzo/models/price"
	"go_base_project/app/response"
)

type PricingService struct {
	Data price.DataPrice
	Repo borzo.BorzoOrderRepository
}

func (service PricingService) Do() response.ServiceResponse {

	result, err := service.Repo.Price(service.Data)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", result)
}