package store

type BusinessHour struct {
	Open  string `json:"open,omitempty"`
	Close string `json:"close,omitempty"`
}

type Location struct {
	StreetNumber string  `json:"street_number"`
	StreetName   string  `json:"street_name"`
	CityName     string  `json:"city_name,omitempty"`
	StateName    string  `json:"state_name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Reference    string  `json:"reference,omitempty"`
}

type CreateRequest struct {
	BusinessHours map[string][]BusinessHour `json:"business_hours,omitempty"`
	ExternalID    string                    `json:"external_id,omitempty"`
	Location      Location                  `json:"location,omitempty"`
	Name          string                    `json:"name,omitempty"`
}
