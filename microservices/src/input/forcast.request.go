package input

type GeoTarget struct {
	Included []CountryTarget `json:"included"`
}

type CountryTarget struct {
	TargetType  string `json:"target_type"`
	TargetID    int    `json:"target_id"`
	Name        string `json:"name"`
	Age        int `json:"age"`
	CountryCode string `json:"country_code"`
}

type DeviceType struct {
	Included []int `json:"included"`
}

type InventoryURL struct {
	Included []string `json:"included"`
}

type ForecastRequest struct {
	GeoTarget    GeoTarget    `json:"geo_target"`
	DeviceType   DeviceType   `json:"device_type"`
	InventoryURL InventoryURL  `json:"inventory_url"`
}
