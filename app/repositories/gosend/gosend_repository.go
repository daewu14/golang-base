package gosend

import (
	"encoding/json"
	"errors"
	"go_base_project/app/repositories/gosend/models/booking"
	"go_base_project/app/repositories/gosend/models/response"
	"go_base_project/app/services/gosend/base"
	"io"
	"strings"
)

type GosendRepoInterface interface {
	Booking(data booking.BookingData) (response.ResponseCreateBookingData, error)
	Pricing()
	CancelBooking()
	StatusBooking()
}

type GosendRepository struct {
}

func (repo GosendRepository) Booking(data booking.BookingData) (response.ResponseCreateBookingData, error) {

	call, err := base.GosendApi{}.Post("gokilat/v10/booking").Bodys(data).Call()

	if err != nil {
		return response.ResponseCreateBookingData{}, err
	}

	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return response.ResponseCreateBookingData{}, errRa
	}

	var responseCreateBookingData response.ResponseCreateBookingData
	errUm := json.Unmarshal(result, &responseCreateBookingData)
	if errUm != nil {
		return response.ResponseCreateBookingData{}, errUm
	}

	/**
	gosend will return array erros if our data was invalid
	convert to string the array then pass it as string error
	*/
	if responseCreateBookingData.Errors != nil {
		joinErrorStsring := strings.Join(responseCreateBookingData.Errors, " ")
		return response.ResponseCreateBookingData{}, errors.New(joinErrorStsring)
	}

	return responseCreateBookingData, nil
}
