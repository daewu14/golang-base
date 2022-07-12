package dto

type PricingDTO struct {
	Origin struct {
		Lat     float64 `json:"lat"`
		Long    float64 `json:"long"`
		Address string  `json:"address"`
	} `json:"origin"`
	Destination struct {
		Lat     float64 `json:"lat"`
		Long    float64 `json:"long"`
		Address string  `json:"address"`
	} `json:"destination"`
	Weight    int      `json:"weight"`
	ItemPrice int      `json:"item_price"`
	Service   []string `json:"service"`
}
