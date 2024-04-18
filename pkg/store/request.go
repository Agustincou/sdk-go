package cardtoken

type BusinessHour struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

type Location struct {
	StreetNumber string  `json:"street_number"`
	StreetName   string  `json:"street_name"`
	CityName     string  `json:"city_name"`
	StateName    string  `json:"state_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Reference    string  `json:"reference"`
}

type CreateRequest struct {
	BusinessHours map[string][]BusinessHour `json:"business_hours"`
	ExternalID    string                    `json:"external_id"`
	Location      Location                  `json:"location"`
	Name          string                    `json:"name"`
}


