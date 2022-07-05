package order

type OrderData struct {
	Type            string   `json:"type"`
	Matter          string   `json:"matter"`
	InsuranceAmount int      `json:"insurance_amount"`
	TotalWeightKg   int      `json:"total_weight_kg"`
	Points          []Points `json:"points"`
}
