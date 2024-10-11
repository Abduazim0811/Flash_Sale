package payment

type PaymentRequest struct {
    OrderID string  `json:"order_id"`
    UserID  string  `json:"user_id"`
    Amount  float32 `json:"amount"`
}

type GetPaymentRequest struct {
    PaymentID string `json:"payment_id"`
}

type Payment struct {
    ID      string  `json:"id"`
    OrderID string  `json:"order_id"`
    UserID  string  `json:"user_id"`
    Amount  float32 `json:"amount"`
    Status  string  `json:"status"`
	CreatedAt string `json:"created_at"`
}
