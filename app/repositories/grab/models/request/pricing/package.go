package pricing

type Package struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Dimensions  struct {
		Height int `json:"height"`
		Width  int `json:"width"`
		Depth  int `json:"depth"`
		Weight int `json:"weight"`
	} `json:"dimensions"`
}
