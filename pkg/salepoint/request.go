package cardtoken

type CreateRequest struct {
	Category        int    `json:"category"`
	ExternalID      string `json:"external_id"`
	ExternalStoreID string `json:"external_store_id"`
	FixedAmount     bool   `json:"fixed_amount"`
	Name            string `json:"name"`
	StoreID         int    `json:"store_id"`
}
