package entity

// ShippingZone shipping zone properties
type ShippingZone struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}
