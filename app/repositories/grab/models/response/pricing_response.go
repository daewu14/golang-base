package response

import "time"

type PricingResponse struct {
	Quotes []struct {
		Service struct {
			Id   int    `json:"id"`
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"service"`
		Currency struct {
			Code     string `json:"code"`
			Symbol   string `json:"symbol"`
			Exponent int    `json:"exponent"`
		} `json:"currency"`
		Amount            int `json:"amount"`
		EstimatedTimeline struct {
			Pickup  time.Time `json:"pickup"`
			Dropoff time.Time `json:"dropoff"`
		} `json:"estimatedTimeline"`
		Distance int `json:"distance"`
	} `json:"quotes"`
	Origin struct {
		Address     string `json:"address"`
		CityCode    string `json:"cityCode"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"origin"`
	Destination struct {
		Address     string `json:"address"`
		CityCode    string `json:"cityCode"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"destination"`
	Packages []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
		Price       int    `json:"price"`
		Dimensions  struct {
		} `json:"dimensions"`
	} `json:"packages"`
}
