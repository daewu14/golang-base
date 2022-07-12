package instant

import (
	"encoding/json"
	"fmt"
	"go_base_project/app/models/incoming_request/instant_service/dto"
	borzo2 "go_base_project/app/repositories/borzo"
	"go_base_project/app/repositories/borzo/models/price"
	"go_base_project/app/repositories/borzo/models/responses"
	"go_base_project/app/repositories/gosend/models/pricing"
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
	channel := make(chan responses.ResponseDataPricing)

	borzoPricingData := contructBorzoPricingData(service.data)
	getBorzoResponse(channel, borzoPricingData)
	fmt.Println("borzo pricing result")

	/*
		- Gosend process
	*/
	gosendPricingData := constructGosendPricingData(service.data)

	gosendPricingResult := gosend.GosendPricingService{gosendPricingData}.Do()

	fmt.Println("gosend pricing result", gosendPricingResult)

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

func getBorzoResponse(jancok chan responses.ResponseDataPricing, dataPrice price.DataPrice) {
	fmt.Println("cok", dataPrice.Type)
	result := borzo.PricingService{dataPrice, borzo2.BorzoOrderRepository{}}.Do()

	var pricingResponse responses.ResponseDataPricing
	pricingResponse, _ = json.Marshal(result.Data, &pricingResponse)
	// sampi disini
	fmt.Println("resultnya", result.Data)
}
