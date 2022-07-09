package gosend

import (
	"go_base_project/app/repositories/gosend"
	"go_base_project/app/response"
)

type GosendStatusBookingService struct {
	orderID string
	repo    gosend.GosendRepository
}

func (service GosendStatusBookingService) Do() response.ServiceResponse {

	result, err := service.repo.StatusBooking(service.orderID)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", result)
}
