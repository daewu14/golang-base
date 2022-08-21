package pricing

type PricingDTOResponse struct {
	Insurance struct {
		Apply bool `json:"apply"`
	} `json:"insurance"`
	Costs []Costs `json:"costs"`
}

type Costs struct {
	Service     string `json:"service"`
	ServiceType string `json:"service_type"`
	Eta         string `json:"eta"`
	Price       Price  `json:"price"`
}

type Price struct {
	AdminFee     string `json:"admin_fee"`
	InsuranceFee string `json:"insurance_fee"`
	ShippingCost int    `json:"shipping_cost"`
	TotalPrice   int    `json:"total_price"`
}