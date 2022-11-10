package pricing

type PricingRequest struct {
	Packages    []Package   `json:"packages"`
	Origin      Origin      `json:"origin"`
	Destination Destination `json:"destination"`
}
