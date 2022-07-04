package borzo

import (
	"fmt"
	"go_base_project/app/repositories/borzo"
	"testing"
)

func TestShowOrderService_Do(t *testing.T) {

	service := ShowOrderService{borzo.BorzoOrderRepository{}}.Do()

	fmt.Println("service result", service)
}
