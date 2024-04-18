package qr

type CreateResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QRData         string `json:"qr_data"`
}
