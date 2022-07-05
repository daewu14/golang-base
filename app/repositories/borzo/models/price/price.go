package price

type DataPrice struct {
	Type          string   `json:"type"`
	TotalWeightKg int      `json:"total_weight_kg"`
	Points        []Points `json:"points"`
}
