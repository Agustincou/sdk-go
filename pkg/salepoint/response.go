package store

import "time"

type QR struct {
	Image         string `json:"image"`
	TemplateDoc   string `json:"template_document"`
	TemplateImage string `json:"template_image"`
}

type QRCode struct {
	Site   string `json:"site"`
	QRCode string `json:"qr_code"`
}

type CreateResponse struct {
	ID              int       `json:"id"`
	QR              QR        `json:"qr"`
	Status          string    `json:"status"`
	DateCreated     time.Time `json:"date_created"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	UUID            string    `json:"uuid"`
	UserID          int       `json:"user_id"`
	Name            string    `json:"name"`
	FixedAmount     bool      `json:"fixed_amount"`
	Category        int       `json:"category"`
	StoreID         string    `json:"store_id"`
	ExternalStoreID string    `json:"external_store_id"`
	ExternalID      string    `json:"external_id"`
	Site            string    `json:"site"`
	QRCode          QRCode    `json:"qr_code"`
}

type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type QrCode struct {
	Site   string `json:"site"`
	QRCode string `json:"qr_code"`
}

type SearchResponseResult struct {
	UserID          int       `json:"user_id"`
	Name            string    `json:"name"`
	FixedAmount     bool      `json:"fixed_amount"`
	Category        int       `json:"category"`
	StoreID         string    `json:"store_id"`
	ExternalID      string    `json:"external_id"`
	ID              int       `json:"id"`
	QR              QR        `json:"qr"`
	DateCreated     time.Time `json:"date_created"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	ExternalStoreID string    `json:"external_store_id"`
	QRCode          string    `json:"qr_code"`
}

type SearchResponse struct {
	Paging  Paging                 `json:"paging"`
	Results []SearchResponseResult `json:"results"`
}
