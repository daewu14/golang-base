package gosend

import (
	"go_base_project/app/repositories/gosend"
	"go_base_project/app/repositories/gosend/models/booking"
	"go_base_project/app/response"
)

type GosendBookingService struct {
	Data booking.BookingData
	Repo gosend.GosendRepository
}

func (service GosendBookingService) Do() response.ServiceResponse {

	result, err := service.Repo.Booking(service.Data)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("OK!", result)

}
