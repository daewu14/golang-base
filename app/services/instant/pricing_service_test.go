package instant

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/models/incoming_request/instant_service/dto"
	"testing"
)

func TestPricingService_Do(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return
	}

	var data dto.PricingDTO
	data.Origin.Lat = -6.3036354
	data.Origin.Long = 106.6947153
	data.Origin.Address = "origin address"
	data.Destination.Lat = -6.285331
	data.Destination.Long = 106.8387787
	data.Destination.Address = "destination address"
	data.Weight = 1000
	data.ItemPrice = 20000
	service := []string{"gosend"}
	data.Service = service

	result := PricingService{Data: data}.Do()

	fmt.Println("result", result)

}

func TestOkeTest(t *testing.T) {
	println("okee")
}
