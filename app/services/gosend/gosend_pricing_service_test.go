package gosend

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/gosend/models/pricing"
	"testing"
)

func TestGosendPricingService_Do(t *testing.T) {

	godotenv.Load("../../../.env")

	var params pricing.PricingData
	params.Origin = "-6.2739117,106.8119382"
	params.Destination = "-6.2339117,106.8219382"

	result := GosendPricingService{params}.Do()

	fmt.Println("TestGosendPricingService_Do result", result)
}
