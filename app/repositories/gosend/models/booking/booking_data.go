package booking

type BookingData struct {
	PaymentType        int     `json:"paymentType"`
	CollectionLocation string  `json:"collection_location"`
	ShipmentMethod     string  `json:"shipment_method"`
	Routes             []Route `json:"routes"`
}
