package store

import "time"

type CreateResponse struct {
	ID            string                    `json:"id"`
	Name          string                    `json:"name"`
	DateCreation  time.Time                 `json:"date_creation"`
	BusinessHours map[string][]BusinessHour `json:"business_hours"`
	Location      LocationResponse          `json:"location"`
	ExternalID    string                    `json:"external_id"`
}

type LocationResponse struct {
	AddressLine string  `json:"address_line"`
	Reference   string  `json:"reference"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	City        string  `json:"city"`
	StateID     string  `json:"state_id"`
}

type SearchResponseResult struct {
	ID            string                    `json:"id"`
	Name          string                    `json:"name"`
	DateCreation  time.Time                 `json:"date_creation"`
	BusinessHours map[string][]BusinessHour `json:"business_hours"`
	Location      LocationResponse          `json:"location"`
	ExternalID    string                    `json:"external_id"`
}

type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type SearchResponse struct {
	Paging  Paging                 `json:"paging"`
	Results []SearchResponseResult `json:"results"`
}
