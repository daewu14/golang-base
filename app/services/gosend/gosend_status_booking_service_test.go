package gosend

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/gosend"
	"testing"
)

func TestGosendStatusBookingService_Do(t *testing.T) {
	godotenv.Load("../../../.env")

	orderID := "GK-11-1478050"

	result := GosendStatusBookingService{orderID: orderID, repo: gosend.GosendRepository{}}.Do()

	fmt.Println("TestGosendStatusBookingService_Do result ", result)
}
