package repository

import (
	"context"
	"order_service/internal/entity/order"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req order.CreateOrderRequest)(*order.CreateOrderResponse, error)
	GetOrder(ctx context.Context, id string) (*order.Order, error)
	GetUserOrders(ctx context.Context, userID string) (*[]order.Order, error)
	GetAllOrders(ctx context.Context) (*[]order.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) (*order.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
}
