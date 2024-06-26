package merchantorder

import "github.com/mercadopago/sdk-go/pkg/date"

// UpdateRequest represents merchant order update.
type UpdateRequest struct {
	Collector *CollectorRequest   `json:"collector,omitempty"`
	Payer     *PayerRequest       `json:"payer,omitempty"`
	Items     []ItemUpdateRequest `json:"items,omitempty"`
	Shipments []ShipmentRequest   `json:"shipments,omitempty"`

	PreferenceID      string `json:"preference_id,omitempty"`
	ApplicationID     string `json:"application_id,omitempty"`
	SiteID            string `json:"site_id,omitempty"`
	NotificationURL   string `json:"notification_url,omitempty"`
	AdditionalInfo    string `json:"additional_info,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Marketplace       string `json:"marketplace,omitempty"`
	Version           int    `json:"version,omitempty"`
	SponsorID         int    `json:"sponsor_id,omitempty"`
}

// ShipmentRequest represents shipment information.
type ShipmentRequest struct {
	ReceiverAddress  *ReceiverAddressRequest `json:"receiver_address,omitempty"`
	ShippingOption   *ShippingOptionRequest  `json:"shipping_option,omitempty"`
	DateCreated      *date.ApiTimeFormat     `json:"date_created,omitempty"`
	LastModified     *date.ApiTimeFormat     `json:"last_modified,omitempty"`
	DateFirstPrinted *date.ApiTimeFormat     `json:"date_first_printed,omitempty"`

	ShippingType      string           `json:"shipping_type,omitempty"`
	ShippingMode      string           `json:"shipping_mode,omitempty"`
	PickingType       string           `json:"picking_type,omitempty"`
	Status            string           `json:"status,omitempty"`
	ShippingSubstatus string           `json:"shipping_substatus,omitempty"`
	ServiceID         string           `json:"service_id,omitempty"`
	ID                int              `json:"id,omitempty"`
	SenderID          int              `json:"sender_id,omitempty"`
	ReceiverID        int              `json:"receiver_id,omitempty"`
	Items             []map[string]any `json:"items,omitempty"`
}

// ReceiverAddressRequest represents receiver address information.
type ReceiverAddressRequest struct {
	City    *CityRequest    `json:"city,omitempty"`
	State   *StateRequest   `json:"state,omitempty"`
	Country *CountryRequest `json:"country,omitempty"`

	AddressLine  string `json:"address_line,omitempty"`
	Apartment    string `json:"apartment,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Contact      string `json:"contact,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	Floor        string `json:"floor,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	ID           int    `json:"id,omitempty"`
}

// ShippingOptionRequest represents shipping option information.
type ShippingOptionRequest struct {
	EstimatedDelivery *EstimatedDeliveryRequest `json:"estimated_delivery,omitempty"`
	Speed             *SpeedRequest             `json:"speed,omitempty"`

	Name             string  `json:"name,omitempty"`
	CurrencyID       string  `json:"currency_id,omitempty"`
	Cost             float64 `json:"cost,omitempty"`
	ListCost         float64 `json:"list_cost,omitempty"`
	ShippingMethodID int     `json:"shipping_method_id,omitempty"`
	ID               int     `json:"id,omitempty"`
}

// CityRequest represents city information.
type CityRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// StateRequest represents state information.
type StateRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// CountryRequest represents country information.
type CountryRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// EstimatedDeliveryRequest represents estimated delivery information.
type EstimatedDeliveryRequest struct {
	Date *date.ApiTimeFormat `json:"date,omitempty"`

	TimeFrom string `json:"time_from,omitempty"`
	TimeTo   string `json:"time_to,omitempty"`
}

// SpeedRequest   represents shipping speed information.
type SpeedRequest struct {
	Handling int `json:"handling,omitempty"`
	Shipping int `json:"shipping,omitempty"`
}

// ItemUpdateRequest represents item information.
type ItemUpdateRequest struct {
	ID       string `json:"id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
