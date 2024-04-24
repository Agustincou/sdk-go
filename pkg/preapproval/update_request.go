package preapproval

import "github.com/mercadopago/sdk-go/pkg/date"

// UpdateRequest represents a request for updating a pre approval.
type UpdateRequest struct {
	AutoRecurring *AutoRecurringUpdateRequest `json:"auto_recurring,omitempty"`

	CardTokenID       string `json:"card_token_id,omitempty"`
	PayerEmail        string `json:"payer_email,omitempty"`
	BackURL           string `json:"back_url,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Status            string `json:"status,omitempty"`
}

// AutoRecurringUpdateRequest represents the recurrence settings.
type AutoRecurringUpdateRequest struct {
	StartDate *date.ApiTimeFormat `json:"start_date,omitempty"`
	EndDate   *date.ApiTimeFormat `json:"end_date,omitempty"`

	TransactionAmount float64 `json:"transaction_amount,omitempty"`
}
