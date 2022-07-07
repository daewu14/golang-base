package gosend

import (
	"go_base_project/app/repositories/gosend"
	"go_base_project/app/repositories/gosend/models/pricing"
	"go_base_project/app/response"
)

type GosendPricingService struct {
	Data pricing.PricingData
}

func (service GosendPricingService) Do() response.ServiceResponse {

	result, err := gosend.GosendRepository{}.Pricing(service.Data)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", result)
}
