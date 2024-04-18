package qr

import "time"

type Item struct {
	SKUNumber   string  `json:"sku_number,omitempty"`
	Category    string  `json:"category,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	UnitMeasure string  `json:"unit_measure,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
}

type CreateRequest struct {
	CashOut struct {
		Amount int `json:"amount,omitempty"`
	} `json:"cash_out,omitempty"`
	Description       string     `json:"description,omitempty"`
	ExternalReference string     `json:"external_reference,omitempty"`
	ExpirationDate    *time.Time `json:"expiration_date,omitempty"`
	Items             []Item     `json:"items,omitempty"`
	NotificationURL   string     `json:"notification_url,omitempty"`
	Title             string     `json:"title,omitempty"`
	TotalAmount       float64    `json:"total_amount,omitempty"`
}
