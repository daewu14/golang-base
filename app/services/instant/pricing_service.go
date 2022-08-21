package instant

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go_base_project/app/models/incoming_request/instant_service/dto"
	pricing2 "go_base_project/app/models/instant/dto/response/pricing"
	borzo2 "go_base_project/app/repositories/borzo"
	"go_base_project/app/repositories/borzo/models/price"
	"go_base_project/app/repositories/borzo/models/responses"
	"go_base_project/app/repositories/gosend/models/pricing"
	response2 "go_base_project/app/repositories/gosend/models/response"
	"go_base_project/app/response"
	"go_base_project/app/services/borzo"
	"go_base_project/app/services/gosend"
	"strconv"
)

type PricingService struct {
	Data dto.PricingDTO
}

func (service PricingService) Do() response.ServiceResponse {

	/*
		- Borzo process
	*/
	borzoChannel := make(chan responses.ResponseDataPricing)
	borzoPricingData := contructBorzoPricingData(service.Data)
	go func() {
		getBorzoResponse(borzoChannel, borzoPricingData)
		close(borzoChannel)
	}()

	/*
		- Gosend process
	*/
	gosendChannel := make(chan response2.ResponsePricingData)
	gosendPricingData := constructGosendPricingData(service.Data)
	go func() {
		getGosendResponse(gosendChannel, gosendPricingData)
		close(gosendChannel)
	}()

	// defined cost for response
	var costs []pricing2.Costs

	golangConstructedData, err := constructPricingData(<-gosendChannel, "gosend", costs)
	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	borzoConstructedData, err := constructPricingData(<-borzoChannel, "borzo", golangConstructedData)
	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success("ok", borzoConstructedData)
}

func contructBorzoPricingData(data dto.PricingDTO) price.DataPrice {
	var borzoPricigData price.DataPrice
	borzoPricigData.Type = "standard"
	borzoPricigData.TotalWeightKg = data.Weight

	var senderContact price.Contact
	var recipentContact price.Contact

	var senderPoint price.Points
	senderPoint.Address = data.Origin.Address
	senderPoint.ContactPerson = senderContact

	var recipentPoint price.Points
	recipentPoint.Address = data.Destination.Address
	recipentPoint.ContactPerson = recipentContact

	borzoPricigData.Points = append(borzoPricigData.Points, senderPoint)
	borzoPricigData.Points = append(borzoPricigData.Points, recipentPoint)

	return borzoPricigData
}

func constructGosendPricingData(data dto.PricingDTO) pricing.PricingData {

	var pricingData pricing.PricingData

	pricingData.Origin = fmt.Sprintf("%f", data.Origin.Lat) + "," + fmt.Sprintf("%f", data.Origin.Long)
	pricingData.Destination = fmt.Sprintf("%f", data.Destination.Lat) + "," + fmt.Sprintf("%f", data.Destination.Long)

	return pricingData
}

func getBorzoResponse(response chan responses.ResponseDataPricing, dataPrice price.DataPrice) {

	var responseDataPricing responses.ResponseDataPricing
	result := borzo.PricingService{dataPrice, borzo2.BorzoOrderRepository{}}.Do()

	errMap := mapstructure.Decode(result.Data, &responseDataPricing)
	if errMap != nil {
		println("something went wrong ! ", errMap.Error())
		return
	}

	response <- responseDataPricing

	return
}

func getGosendResponse(response chan response2.ResponsePricingData, data pricing.PricingData) {

	var responseDataPricing response2.ResponsePricingData
	gosendPricingResult := gosend.GosendPricingService{Data: data}.Do()

	if gosendPricingResult.Status != true {
		println("something went wrong !")
		response <- responseDataPricing
		return
	}

	errMap := mapstructure.Decode(gosendPricingResult.Data, &responseDataPricing)
	if errMap != nil {
		println("something went wrong ! ", errMap.Error())
		response <- responseDataPricing
		return
	}

	response <- responseDataPricing
	return
}

func constructPricingData(data interface{}, service string, cost []pricing2.Costs) ([]pricing2.Costs, error) {

	// defined temp data for cost
	var tempCost pricing2.Costs

	switch service {
	case "gosend":

		// mapping interface to gosend response pricing data
		var responsePricingData response2.ResponsePricingData
		err := mapstructure.Decode(data, &responsePricingData)
		if err != nil {
			return cost, err
		}

		//instant data
		tempCost.Price.TotalPrice = responsePricingData.Instant.Price.TotalPrice
		tempCost.Service = "gosend"
		tempCost.ServiceType = "instant"
		tempCost.Eta = responsePricingData.Instant.ShipmentMethodDescription
		cost = append(cost, tempCost)

		// sameday data
		tempCost.Price.TotalPrice = responsePricingData.SameDay.Price.TotalPrice
		tempCost.Service = "gosend"
		tempCost.ServiceType = "sameday"
		tempCost.Eta = responsePricingData.SameDay.ShipmentMethodDescription
		cost = append(cost, tempCost)

	case "borzo":
		// mapping interface to borzo response pricing data
		var responseDataPricing responses.ResponseDataPricing
		err := mapstructure.Decode(data, &responseDataPricing)
		if err != nil {
			return cost, err
		}
		tempCost.Price.TotalPrice, _ = strconv.Atoi(responseDataPricing.Order.PaymentAmount)
		tempCost.Service = "borzo"
		tempCost.ServiceType = "sameday"
		cost = append(cost, tempCost)
	}

	return cost, nil
}
