package pricing

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type PricingDTORequest struct {
	Origin      *Origin      `json:"origin"`
	Destination *Destination `json:"destination"`
	Weight      int          `json:"weight"`
	ItemPrice   int          `json:"item_price"`
	Service     []string     `json:"service"`
}

type Destination struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Address string  `json:"address"`
}

type Origin struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Address string  `json:"address"`
}

func (a PricingDTORequest) ValidationPricingDTORequest() error {

	return validation.ValidateStruct(&a, validation.Field(&a.Weight, validation.Required.Error("Weight cannot be empty")),
		validation.Field(&a.ItemPrice, validation.Required.Error("Item price cannot be empty")),
		validation.Field(&a.Service, validation.Required.Error("Service cannot be null !")),
		validation.Field(&a.Destination, validation.Required.Error("destination cannot be empty")),
		validation.Field(&a.Origin, validation.Required.Error("Origin cannot be empty")),
	)
}

func (a Origin) ValidateOrigin() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Lat, validation.Required.Error("Latitude cannot be empty")),
		validation.Field(&a.Long, validation.Required.Error("Longitude cannot be empty")),
		validation.Field(&a.Address, validation.Required.Error("Address cannot be empty")),
	)
}

func (a Destination) ValidationDestination() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Lat, validation.Required.Error("Latitude cannot be empty")),
		validation.Field(&a.Long, validation.Required.Error("Longitude cannot be empty")),
		validation.Field(&a.Address, validation.Required.Error("Address cannot be empty")))
}
