package borzo

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/borzo"
	"testing"
)

func TestCancelOrderService_Do(t *testing.T) {
	godotenv.Load("../../../.env")

	result := CancelOrderService{"13544", borzo.BorzoOrderRepository{}}.Do()

	fmt.Println("TestCancelOrderService_Do result ", result)
}
