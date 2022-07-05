package borzo

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_base_project/app/repositories/borzo/models/order"
	"go_base_project/app/repositories/borzo/models/price"
	"testing"
)

func TestBorzoOrderRepository_Order(t *testing.T) {
	godotenv.Load("../../../.env")

	var param order.OrderData

	var senderContact order.Contact
	senderContact.Name = "gema antika hariadi"
	senderContact.Phone = "0899234234"

	var recipientContact order.Contact
	recipientContact.Name = "si gema"
	recipientContact.Phone = "0899923423423"

	var senderPoint order.Points
	senderPoint.Address = "JL. Jawa, Blok J1 No. 31, Komplek Nusaloka, Tangerang, Rw. Mekar Jaya, Serpong, Kota Tangerang Selatan, Banten 15310"
	senderPoint.ContactPerson = senderContact

	var recipientPoint order.Points
	recipientPoint.Address = "Jl. Raya Ragunan No.39, RT.1/RW.2, Ps. Minggu, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12540"
	recipientPoint.ContactPerson = recipientContact

	param.Points = append(param.Points, senderPoint)
	param.Points = append(param.Points, recipientPoint)

	param.Matter = "document"
	param.TotalWeightKg = 12
	response, err := BorzoOrderRepository{}.Order(param)

	fmt.Println(response, err)
}

func TestBorzoOrderRepository_Price(t *testing.T) {
	godotenv.Load("../../../.env")

	var params price.DataPrice

	params.Type = "standard"
	params.TotalWeightKg = 100

	var senderContact price.Contact
	senderContact.Name = "gema antika hariadi"
	senderContact.Phone = "0899234234"

	var recipientContact price.Contact
	recipientContact.Name = "si gema"
	recipientContact.Phone = "0899923423423"

	var senderPoint price.Points
	senderPoint.Address = "JL. Jawa, Blok J1 No. 31, Komplek Nusaloka, Tangerang, Rw. Mekar Jaya, Serpong, Kota Tangerang Selatan, Banten 15310"
	senderPoint.ContactPerson = senderContact

	var recipientPoint price.Points
	recipientPoint.Address = "Jl. Raya Ragunan No.39, RT.1/RW.2, Ps. Minggu, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12540"
	recipientPoint.ContactPerson = recipientContact

	params.Points = append(params.Points, senderPoint)
	params.Points = append(params.Points, recipientPoint)

	response, err := BorzoOrderRepository{}.Price(params)

	fmt.Println("Amount : ", response.Order.PaymentAmount)
	fmt.Println(response, err)
}
