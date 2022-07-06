package booking

type Route struct {
	OriginName              string          `json:"originName"`
	OriginNote              string          `json:"originNote"`
	OriginContactName       string          `json:"originContactName"`
	OriginContactPhone      string          `json:"originContactPhone"`
	OriginLatLong           string          `json:"originLatLong"`
	OriginAddress           string          `json:"originAddress"`
	DestinationName         string          `json:"destinationName"`
	DestinationNote         string          `json:"destinationNote"`
	DestinationContactName  string          `json:"destinationContactName"`
	DestinationContactPhone string          `json:"destinationContactPhone"`
	DestinationLatLong      string          `json:"destinationLatLong"`
	DestinationAddress      string          `json:"destinationAddress"`
	Item                    string          `json:"item"`
	StoreOrderID            string          `json:"storeOrderId"`
	InsuranceDetails        InsuranceDetail `json:"insuranceDetails"`
}
