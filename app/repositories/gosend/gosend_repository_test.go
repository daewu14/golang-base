package gosend

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/gosend/models/booking"
	"go_base_project/app/repositories/gosend/models/pricing"
	"testing"
)

func TestGosendRepository_Booking(t *testing.T) {
	godotenv.Load("../../../.env")
	var dataBooking booking.BookingData
	dataBooking.PaymentType = 3
	dataBooking.ShipmentMethod = "sameday"
	dataBooking.CollectionLocation = "pickup"

	var route booking.Route
	route.OriginName = "si daewu origin"
	route.OriginNote = "si daewu origin note"
	route.OriginContactName = "si daewu origin contact name"
	route.OriginContactPhone = "089923894234"
	route.OriginLatLong = "-6.2739117,106.8119382"
	route.OriginAddress = "Jalan palagan tentara pelajar"
	route.DestinationName = "gema ganteng"
	route.DestinationNote = "note si cakep destination"
	route.DestinationContactName = "gema ganteng"
	route.DestinationContactPhone = "089782349234"
	route.DestinationLatLong = "-6.2339117,106.8219382"
	route.DestinationAddress = "mosok jalan jauh banget"
	route.Item = "cepatu"

	var insuranceDetail booking.InsuranceDetail
	insuranceDetail.Fee = "500"
	insuranceDetail.Applied = "true"
	insuranceDetail.ProductDescription = "pokoke iku lah"
	insuranceDetail.ProductPrice = "150000000"

	route.InsuranceDetails = insuranceDetail

	dataBooking.Routes = append(dataBooking.Routes, route)

	result, err := GosendRepository{}.Booking(dataBooking)

	fmt.Println("TestGosendRepository_Booking result ", result.Id)
	fmt.Println("TestGosendRepository_Booking error ", err)
}

func TestGosendRepository_Pricing(t *testing.T) {
	godotenv.Load("../../../.env")

	var pricingData pricing.PricingData

	pricingData.Origin = "-6.2739117,106.8119382"
	pricingData.Destination = "-6.2339117,106.8219382"

	result, err := GosendRepository{}.Pricing(pricingData)

	fmt.Println("TestGosendRepository_Pricing result", result)
	fmt.Println("TestGosendRepository_Pricing error", err)
}
