package borzo

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/borzo"
	"testing"
)

func TestShowOrderService_Do(t *testing.T) {
	godotenv.Load("../../../.env")
	service := ShowOrderService{"12504", borzo.BorzoOrderRepository{}}.Do()

	fmt.Println("service result", service)
}
