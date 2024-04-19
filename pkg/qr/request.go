package qr

import "github.com/mercadopago/sdk-go/pkg/date"

type Item struct {
	SKUNumber   string  `json:"sku_number,omitempty"`
	Category    string  `json:"category,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	UnitMeasure string  `json:"unit_measure,omitempty"`
	TotalAmount float64 `json:"total_amount"`
}

type CashOut struct {
	Amount int `json:"amount,omitempty"`
}

type CreateRequest struct {
	CashOut           *CashOut            `json:"cash_out,omitempty"`
	Description       string              `json:"description,omitempty"`
	ExternalReference string              `json:"external_reference,omitempty"`
	ExpirationDate    *date.ApiTimeFormat `json:"expiration_date,omitempty"`
	Items             []Item              `json:"items,omitempty"`
	NotificationURL   string              `json:"notification_url,omitempty"`
	Title             string              `json:"title,omitempty"`
	TotalAmount       float64             `json:"total_amount,omitempty"`
}
