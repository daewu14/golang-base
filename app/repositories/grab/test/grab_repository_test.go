package test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"go_base_project/app/repositories/grab"
	"go_base_project/app/repositories/grab/models"
	"go_base_project/app/repositories/grab/models/request"
	"go_base_project/app/repositories/grab/models/request/pricing"
	"go_base_project/app/repositories/grab/models/response"
	grab2 "go_base_project/app/services/grab"
	"go_base_project/app/services/grab/base"
	"os"
	"testing"
)

func TestGrabRepository_GetToken(t *testing.T) {

	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("the err", err)
	}

	// defined data
	var getTokenRequest request.GetTokenRequest
	getTokenRequest.ClientID = base.GrabConfig{}.ClientID()
	getTokenRequest.GrantType = "client_credentials"
	getTokenRequest.Scope = "grab_express.partner_deliveries"
	getTokenRequest.ClientSecret = base.GrabConfig{}.ClientSecret()

	// call repo
	result, errRepo := grab.GrabRepository{}.GetToken(getTokenRequest)

	if errRepo != nil {
		println("something went wrong ! ", errRepo.Error())
		return
	}

	fmt.Println(result)
	return
}

func TestGrabRepository_Pricing(t *testing.T) {

	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("the err", err)
	}

	// defined data
	var data pricing.PricingRequest

	// defined origin
	var origin pricing.Origin
	origin.Address = "Jl. Palagan Tentara Pelajar No.KM 08, Karang Moko, Sariharjo, Kec. Ngaglik, Kabupaten Sleman, Daerah Istimewa Yogyakarta 55581"
	origin.Coordinates.Latitude = -7.730787
	origin.Coordinates.Longitude = 110.3767362

	// defined destination
	var destination pricing.Destination
	destination.Address = "Magelang St Jl. Gito Gati No.KM 9, Denggung, Tridadi, Sleman, Sleman Regency, Special Region of Yogyakarta 55511"
	destination.Coordinates.Latitude = -7.7201812
	destination.Coordinates.Longitude = 110.3595438

	// defined package
	var pcg pricing.Package
	pcg.Name = "Fish Burger"
	pcg.Description = "Fish Burger with mayonnaise sauce"
	pcg.Quantity = 1
	pcg.Dimensions.Weight = 2000
	pcg.Dimensions.Height = 10
	pcg.Dimensions.Width = 10
	pcg.Dimensions.Depth = 10

	data.Origin = origin
	data.Destination = destination
	data.Packages = append(data.Packages, pcg)

	// call get token service

	var getTokenResponse response.GetTokenResponse

	service := grab2.GetTokenService{
		grab.GrabRepository{},
	}.Do()

	if service.Status != true {
		println("something went wrong ! ", service.Message)
		return
	}

	if errMap := mapstructure.Decode(service.Data, &getTokenResponse); errMap != nil {
		println("something went wrong ! ", errMap.Error())
		return
	}

	// defined credential
	var credential models.Credential
	credential.Bearer = getTokenResponse.AccessToken

	// call repo

	repo, err := grab.GrabRepository{}.Pricing(credential, data)

	if err != nil {
		println("something went wrong ! ", err.Error())
		return
	}

	log.WithFields(log.Fields{
		"response": repo,
	}).Info("GRAB : SIPALING TESTING")

}
