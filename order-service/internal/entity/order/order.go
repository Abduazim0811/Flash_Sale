package order

import "time"

type Order struct {
	OrderId   string      `json:"order_id" bson:"order_id"`
	UserID     string      `json:"user_id" bson:"user_id"`
	Items      []OrderItem `json:"items" bson:"items"`
	TotalPrice float64     `json:"total_price" bson:"total_price"`
	Status     string      `json:"status" bson:"status"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" bson:"updated_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int32   `json:"quantity" bson:"quantity"`
}

type CreateOrderRequest struct {
	OrderId  string	   `json:"order_id" bson:"order_id"`
	UserID string      `json:"user_id" bson:"user_id"`
	Items  []OrderItem `json:"items" bson:"items"`
}

type CreateOrderResponse struct{
	Status     string      `json:"status" bson:"status"`
	OrderId   string      `json:"order_id" bson:"order_id"`
	Message   string		`json:"message" bson:"message"`
	TotalPrice float64     `json:"total_price" bson:"total_price"`
	Product  []UnavailableProduct `json:"prodcut" bson:"product"`
}

type GetOrdersRequest struct {
	UserID string `json:"user_id" bson:"user_id"`
}

type GetOrdersResponse struct {
	Orders []Order `json:"orders" bson:"orders"`
}

type GetOrderRequest struct {
	OrderId string `json:"order_id" bson:"order_id"`
}

type GetOrderResponse struct {
	Order Order `json:"order" bson:"order"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"order_id" bson:"order_id"`
	Status  string `json:"status" bson:"status"`
}

type UpdateOrderStatusResponse struct {
	Order Order `json:"order" bson:"order"`
}

type GetAllOrdersRequest struct{}

type GetAllOrdersResponse struct {
	Orders []Order `json:"orders" bson:"orders"`
}


type UnavailableProduct struct{
	ProductID string  `json:"product_id" bson:"product_id"`
	RequestedQuantity int32 `json:"requested_quantity" bson:"requested_quantity"`
	AvailableQuantity int32 `json:"available_quantity" bson:"available_quantity"`
}