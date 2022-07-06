package gosend

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/gosend"
	"go_base_project/app/repositories/gosend/models/booking"
	"testing"
)

func TestGosendBookingService_Do(t *testing.T) {

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

	result := GosendBookingService{dataBooking, gosend.GosendRepository{}}.Do()

	fmt.Println("TestGosendBookingService_Do result ", result)

}
