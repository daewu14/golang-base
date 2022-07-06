package response

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
