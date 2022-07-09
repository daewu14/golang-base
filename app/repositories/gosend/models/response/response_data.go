package response

import "time"

type ResponseCreateBookingData struct {
	Id             int         `json:"id"`
	OrderNo        string      `json:"orderNo"`
	BookingType    string      `json:"bookingType"`
	StoreOrderId   string      `json:"storeOrderId"`
	ErrorMessage   interface{} `json:"errorMessage"`
	Prebook        bool        `json:"prebook"`
	PrebookMessage interface{} `json:"prebookMessage"`
	Errors         []string    `json:"errors"`
}

type ResponsePricingData struct {
	Instant struct {
		ShipmentMethod            string  `json:"shipment_method"`
		ShipmentMethodDescription string  `json:"shipment_method_description"`
		Serviceable               bool    `json:"serviceable"`
		Active                    bool    `json:"active"`
		Distance                  float64 `json:"distance"`
		RoutePolyline             string  `json:"route_polyline"`
		Price                     struct {
			TotalPrice                 int         `json:"total_price"`
			GoPayTotalPrice            int         `json:"go_pay_total_price"`
			GoPayDiscount              int         `json:"go_pay_discount"`
			GoPayMessage               interface{} `json:"go_pay_message"`
			VoucherDiscount            int         `json:"voucher_discount"`
			VoucherGoPayDiscount       int         `json:"voucher_go_pay_discount"`
			GoPayDiscountWithVoucher   int         `json:"go_pay_discount_with_voucher"`
			VoucherMessage             string      `json:"voucher_message"`
			VoucherMonetaryValue       int         `json:"voucher_monetary_value"`
			TotalPriceWithVoucher      int         `json:"total_price_with_voucher"`
			GoPayTotalPriceWithVoucher int         `json:"go_pay_total_price_with_voucher"`
		} `json:"price"`
	} `json:"Instant"`
	SameDay struct {
		ShipmentMethod            string  `json:"shipment_method"`
		ShipmentMethodDescription string  `json:"shipment_method_description"`
		Serviceable               bool    `json:"serviceable"`
		Active                    bool    `json:"active"`
		Distance                  float64 `json:"distance"`
		RoutePolyline             string  `json:"route_polyline"`
		Price                     struct {
			TotalPrice                 int         `json:"total_price"`
			GoPayTotalPrice            int         `json:"go_pay_total_price"`
			GoPayDiscount              int         `json:"go_pay_discount"`
			GoPayMessage               interface{} `json:"go_pay_message"`
			VoucherDiscount            int         `json:"voucher_discount"`
			VoucherGoPayDiscount       int         `json:"voucher_go_pay_discount"`
			GoPayDiscountWithVoucher   int         `json:"go_pay_discount_with_voucher"`
			VoucherMessage             string      `json:"voucher_message"`
			VoucherMonetaryValue       int         `json:"voucher_monetary_value"`
			TotalPriceWithVoucher      int         `json:"total_price_with_voucher"`
			GoPayTotalPriceWithVoucher int         `json:"go_pay_total_price_with_voucher"`
		} `json:"price"`
	} `json:"SameDay"`
}

type ResponseStatusBookingData struct {
	ID                  int         `json:"id"`
	OrderNo             string      `json:"orderNo"`
	Status              string      `json:"status"`
	DriverID            interface{} `json:"driverId"`
	DriverName          interface{} `json:"driverName"`
	DriverPhone         interface{} `json:"driverPhone"`
	DriverPhoto         interface{} `json:"driverPhoto"`
	VehicleNumber       interface{} `json:"vehicleNumber"`
	TotalDiscount       int         `json:"totalDiscount"`
	TotalPrice          int         `json:"totalPrice"`
	ReceiverName        interface{} `json:"receiverName"`
	OrderCreatedTime    time.Time   `json:"orderCreatedTime"`
	OrderDispatchTime   interface{} `json:"orderDispatchTime"`
	OrderArrivalTime    interface{} `json:"orderArrivalTime"`
	OrderClosedTime     interface{} `json:"orderClosedTime"`
	SellerAddressName   string      `json:"sellerAddressName"`
	SellerAddressDetail string      `json:"sellerAddressDetail"`
	BuyerAddressName    string      `json:"buyerAddressName"`
	BuyerAddressDetail  string      `json:"buyerAddressDetail"`
	BookingType         string      `json:"bookingType"`
	CancelDescription   interface{} `json:"cancelDescription"`
	StoreOrderID        string      `json:"storeOrderId"`
	LiveTrackingURL     string      `json:"liveTrackingUrl"`
	InsuranceDetails    struct {
		Applied            string `json:"applied"`
		Fee                string `json:"fee"`
		ProductDescription string `json:"product_description"`
		ProductPrice       string `json:"product_price"`
	} `json:"insuranceDetails"`
	Proofs struct {
		Pop interface{} `json:"pop"`
		Pod interface{} `json:"pod"`
	} `json:"proofs"`
	BookingStatus string `json:"bookingStatus"`
}
