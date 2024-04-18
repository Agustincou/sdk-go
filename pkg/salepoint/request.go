package salepoint

type CreateRequest struct {
	Category        int    `json:"category,omitempty"`
	ExternalID      string `json:"external_id,omitempty"`
	ExternalStoreID string `json:"external_store_id,omitempty"`
	FixedAmount     bool   `json:"fixed_amount,omitempty"`
	Name            string `json:"name,omitempty"`
	StoreID         int    `json:"store_id,omitempty"`
}

// No agregar "omitempty" porque rompe la construcción de los queryParams
type SearchRequest struct {
	UserID          int    `json:"user_id"`
	Category        int    `json:"category"`
	StoreID         string `json:"store_id"`
	ExternalID      string `json:"external_id"`
	ExternalStoreID string `json:"external_store_id"`
}
