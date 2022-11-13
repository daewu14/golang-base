package grab

import (
	"github.com/mitchellh/mapstructure"
	"go_base_project/app/repositories/grab"
	"go_base_project/app/repositories/grab/models"
	request "go_base_project/app/repositories/grab/models/request/pricing"
	response2 "go_base_project/app/repositories/grab/models/response"
	"go_base_project/app/response"
)

type PricingService struct {
	Data request.PricingRequest
	Repo grab.GrabRepoInterface
}

func (s PricingService) Do() response.ServiceResponse {

	// get access token
	accessTokenService := GetTokenService{Repo: grab.GrabRepository{}}.Do()
	if !accessTokenService.Status {
		return response.Service().Error(accessTokenService.Message, accessTokenService.Data)
	}

	// mapping access token response
	var getTokenResponse response2.GetTokenResponse
	if errMap := mapstructure.Decode(accessTokenService.Data, &getTokenResponse); errMap != nil {
		return response.Service().Error(errMap.Error(), nil)
	}

	// define credential
	var credential models.Credential
	credential.Bearer = getTokenResponse.AccessToken

	repo, errRepo := s.Repo.Pricing(credential, s.Data)
	if errRepo != nil {
		return response.Service().Error(errRepo.Error(), nil)
	}

	return response.Service().Success("OK!", repo)
}
