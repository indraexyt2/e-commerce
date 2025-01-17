package models

type PaymentInitiatePayload struct {
	UserID     int     `json:"user_id"`
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}
