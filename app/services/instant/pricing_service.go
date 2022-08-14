package instant

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go_base_project/app/models/incoming_request/instant_service/dto"
	borzo2 "go_base_project/app/repositories/borzo"
	"go_base_project/app/repositories/borzo/models/price"
	"go_base_project/app/repositories/borzo/models/responses"
	"go_base_project/app/repositories/gosend/models/pricing"
	response2 "go_base_project/app/repositories/gosend/models/response"
	"go_base_project/app/response"
	"go_base_project/app/services/borzo"
	"go_base_project/app/services/gosend"
)

type PricingService struct {
	data dto.PricingDTO
}

func (service PricingService) Do() response.ServiceResponse {

	/*
		- Borzo process
	*/
	borzoChannel := make(chan responses.ResponseDataPricing)

	borzoPricingData := contructBorzoPricingData(service.data)
	go getBorzoResponse(borzoChannel, borzoPricingData)

	/*
		- Gosend process
	*/
	gosendChannel := make(chan response2.ResponsePricingData)

	gosendPricingData := constructGosendPricingData(service.data)
	go getGosendResponse(gosendChannel, gosendPricingData)

	fmt.Println("gosend channel result ", <-gosendChannel)
	fmt.Println("borzo channel result ", <-borzoChannel)

	return response.Service().Success("ok", nil)
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
