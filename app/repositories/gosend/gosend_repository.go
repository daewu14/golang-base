package gosend

import (
	"encoding/json"
	"errors"
	"go_base_project/app/repositories/gosend/models/booking"
	"go_base_project/app/repositories/gosend/models/pricing"
	"go_base_project/app/repositories/gosend/models/response"
	"go_base_project/app/services/gosend/base"
	"io"
	"strings"
)

type GosendRepoInterface interface {
	Booking(data booking.BookingData) (response.ResponseCreateBookingData, error)
	Pricing(data pricing.PricingData) (response.ResponsePricingData, error)
	CancelBooking()
	StatusBooking(orderID string) (response.ResponseStatusBookingData, error)
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

func (repo GosendRepository) Pricing(data pricing.PricingData) (response.ResponsePricingData, error) {

	var params = map[string]string{
		"origin":      data.Origin,
		"destination": data.Destination,
	}

	call, err := base.GosendApi{}.Get("gokilat/v10/calculate/price").Params(params).Call()
	if err != nil {
		return response.ResponsePricingData{}, err
	}

	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return response.ResponsePricingData{}, errRa
	}

	var responsePricingData response.ResponsePricingData
	errUm := json.Unmarshal(result, &responsePricingData)
	if errUm != nil {
		return response.ResponsePricingData{}, errUm
	}

	return responsePricingData, nil
}

func (repo GosendRepository) StatusBooking(orderID string) (response.ResponseStatusBookingData, error) {
	endpoint := "gokilat/v10/booking/orderno/" + orderID
	call, err := base.GosendApi{}.Get(endpoint).Call()

	if err != nil {
		return response.ResponseStatusBookingData{}, err
	}

	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return response.ResponseStatusBookingData{}, errRa
	}

	var responseStatusBookingData response.ResponseStatusBookingData
	errUm := json.Unmarshal(result, &responseStatusBookingData)
	if errUm != nil {
		return response.ResponseStatusBookingData{}, errUm
	}

	return responseStatusBookingData, nil

}
