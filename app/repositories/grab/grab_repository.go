package grab

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go_base_project/app/repositories/grab/models"
	"go_base_project/app/repositories/grab/models/request"
	"go_base_project/app/repositories/grab/models/request/pricing"
	"go_base_project/app/repositories/grab/models/response"
	"go_base_project/app/services/grab/base"
	"io"
)

type GrabRepoInterface interface {
	GetToken(request request.GetTokenRequest) (response.GetTokenResponse, error)
	Pricing(credential models.Credential, request pricing.PricingRequest) (response.PricingResponse, error)
}

type GrabRepository struct {
}

func (repo GrabRepository) GetToken(tokenRequest request.GetTokenRequest) (response.GetTokenResponse, error) {

	// defined variable get token response
	var getTokenResponse response.GetTokenResponse

	// call get token api
	call, err := base.GrabApi{}.PostGetAccessToken("grabid/v1/oauth2/token").Bodys(tokenRequest).Call()
	if err != nil {
		return getTokenResponse, err
	}

	// read get token response body
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return getTokenResponse, errRa
	}

	// mapping the response body to get token response variable
	errMap := json.Unmarshal(result, &getTokenResponse)
	if errMap != nil {
		return getTokenResponse, errMap
	}

	log.WithFields(log.Fields{
		"response": getTokenResponse,
	}).Info("GRAB : GET TOKEN")

	return getTokenResponse, nil
}

func (repo GrabRepository) Pricing(credential models.Credential, pricingRequest pricing.PricingRequest) (response.PricingResponse, error) {

	// defined variable pricing response
	var pricingResponse response.PricingResponse

	// call pricing api
	call, err := base.GrabApi{}.Post("v1/deliveries/quotes").Bodys(pricingRequest).AddHeader("Authorization", "Bearer "+credential.Bearer).Call()
	if err != nil {
		return pricingResponse, err
	}

	// read pricing response body
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return pricingResponse, errRa
	}

	// mapping the response body to pricing response variable
	if errMap := json.Unmarshal(result, &pricingResponse); errMap != nil {
		return pricingResponse, errMap
	}

	log.WithFields(log.Fields{
		"response": pricingResponse,
	}).Info("GRAB : PRICING RESPONSE")

	return pricingResponse, nil

}
