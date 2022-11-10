package grab

import (
	"go_base_project/app/repositories/grab"
	"go_base_project/app/repositories/grab/models/request"
	"go_base_project/app/response"
	"go_base_project/app/services/grab/base"
)

type GetTokenService struct {
	Repo grab.GrabRepoInterface
}

func (s GetTokenService) Do() response.ServiceResponse {

	// defined get token request data
	var getTokenRequest request.GetTokenRequest
	getTokenRequest.ClientID = base.GrabConfig{}.ClientID()
	getTokenRequest.GrantType = "client_credentials"
	getTokenRequest.Scope = "grab_express.partner_deliveries"
	getTokenRequest.ClientSecret = base.GrabConfig{}.ClientSecret()

	// call repo
	repo, err := s.Repo.GetToken(getTokenRequest)
	if err != nil {
		return response.ServiceResponse{}.Error(err.Error(), nil)
	}

	if repo.Error != "" {
		return response.ServiceResponse{}.Error(repo.Error+" "+repo.ErrorDescription, nil)
	}

	return response.ServiceResponse{}.Success("success", repo)

}
