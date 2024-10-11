package flashsale

type FlashSaleProduct struct {
	ProductID      string  `json:"product_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	OriginalPrice  float32 `json:"original_price"`
	DiscountPrice  float32 `json:"discount_price"`
	Stock          int32   `json:"stock"`
	IsActive       bool    `json:"is_active"`
}

type FlashSaleEvent struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	StartTime   string             `json:"start_time"`
	EndTime     string             `json:"end_time"`
	Products    []FlashSaleProduct `json:"products"`
	IsActive    bool               `json:"is_active"`
}

type CreateFlashSaleRequest struct {
	FlashSaleId string			   `json:"flashsale_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	StartTime   string             `json:"start_time"`
	EndTime     string             `json:"end_time"`
	Products    []FlashSaleProduct `json:"products"`
}

type CreateFlashSaleResponse struct {
	Event FlashSaleEvent `json:"event"`
}

type UpdateFlashSaleRequest struct {
	FlashSaleId string			   `json:"flashsale_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	StartTime   string             `json:"start_time"`
	EndTime     string             `json:"end_time"`
	Products    []FlashSaleProduct `json:"products"`
}

type UpdateFlashSaleResponse struct {
	Event FlashSaleEvent `json:"event"`
}

type DeleteFlashSaleRequest struct {
	FlashSaleId string			   `json:"flashsale_id"`
}

type DeleteFlashSaleResponse struct {
	Message string `json:"message"`
}

type ListFlashSalesRequest struct{}

type ListFlashSalesResponse struct {
	Events []FlashSaleEvent `json:"events"`
}

type PurchaseProductRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	UserID    string `json:"user_id"`
}

type PurchaseProductResponse struct {
	Message string           `json:"message"`
	Product FlashSaleProduct `json:"product"`
}

type GetActiveFlashSalesRequest struct{}

type GetActiveFlashSalesResponse struct {
	ActiveEvents []FlashSaleEvent `json:"active_events"`
}