package test

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	grab2 "go_base_project/app/repositories/grab"
	"go_base_project/app/repositories/grab/models/request/pricing"
	"go_base_project/app/services/grab"
	"os"
	"testing"
)

func TestPricingService_Do(t *testing.T) {

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
	var pricingRequest pricing.PricingRequest
	pricingRequest.Origin.Address = "ok"

	service := grab.PricingService{
		Data: pricingRequest,
		Repo: grab2.GrabRepository{},
	}.Do()

	if service.Status != true {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("result ", service.Data)
}
