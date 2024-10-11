package service

import (
	"context"
	"fmt"
	"log"
	userclients "order_service/internal/clients/user_clients"
	"order_service/internal/entity/order"
	"order_service/internal/infrastructura/repository"
	"order_service/protos/order_proto"
)

type OrderService struct{
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService{
	return &OrderService{repo: repo}
}

func (o *OrderService) Createorder(ctx context.Context, req *order_proto.CreateOrderRequest) (*order_proto.CreateOrderResponse, error) {
	err := userclients.GetUserId(req.UserId)
	if err != nil {
		log.Println("users not found:", err)
		return nil, fmt.Errorf("users not found: %v", err)
	}
	orderReq := order.CreateOrderRequest{
		UserID: req.UserId,
		Items:  mapProtoItemsToOrder(req.Items),
	}
	res, err := o.repo.CreateOrder(ctx, orderReq)
	return &order_proto.CreateOrderResponse{
		Status: res.Status,
		OrderId: res.OrderId,
		TotalAmount: res.TotalPrice,
		Message: res.Message,
		UnavailableProducts: mapUnavailableProducts(res.Product),
	}, err
}

func (o *OrderService) Getorder(ctx context.Context, orderId string) (*order_proto.GetOrderResponse, error) {

	res, err := o.repo.GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	orderProto := &order_proto.Order{
		Id:         res.OrderId,
		UserId:     res.UserID,
		Items:      mapOrderItemsToProto(res.Items),
		TotalPrice: res.TotalPrice,
		Status:     res.Status,
		CreatedAt:  res.CreatedAt.String(),
		UpdatedAt:  res.UpdatedAt.String(),
	}

	return &order_proto.GetOrderResponse{
		Order: orderProto,
	}, nil
}

func (o *OrderService) GetUserorders(ctx context.Context, userId string) (*order_proto.GetOrdersResponse, error) {
	res, err := o.repo.GetUserOrders(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("get user orders error: %v", err)
	}

	var protoOrders []*order_proto.Order
	for _, order := range *res {
		protoOrders = append(protoOrders, &order_proto.Order{
			Id:         order.OrderId,
			UserId:     order.UserID,
			Items:      mapOrderItemsToProto(order.Items),
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.String(),
			UpdatedAt:  order.UpdatedAt.String(),
		})
	}

	return &order_proto.GetOrdersResponse{
		Orders: protoOrders,
	}, nil
}


func (o *OrderService) GetAllorders(ctx context.Context) (*order_proto.GetAllOrdersResponse, error) {
	res, err := o.repo.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all orders error: %v", err)
	}

	var protoOrders []*order_proto.Order
	for _, order := range *res {
		protoOrders = append(protoOrders, &order_proto.Order{
			Id:         order.OrderId,
			UserId:     order.UserID,
			Items:      mapOrderItemsToProto(order.Items),
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.String(),
			UpdatedAt:  order.UpdatedAt.String(),
		})
	}

	return &order_proto.GetAllOrdersResponse{
		Orders: protoOrders,
	}, nil
}


func (o *OrderService) UpdateorderStatus(ctx context.Context, orderID string, status string) (*order_proto.Order, error) {
	res, err := o.repo.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		return nil, fmt.Errorf("update order status error: %v", err)
	}
	
	orderProto := &order_proto.Order{
		Id:         res.OrderId,
		UserId:     res.UserID,
		Items:      mapOrderItemsToProto(res.Items),
		TotalPrice: res.TotalPrice,
		Status:     res.Status,
		CreatedAt:  res.CreatedAt.String(),
		UpdatedAt:  res.UpdatedAt.String(),
	}

	return orderProto, nil
}


func (o *OrderService) Deleteorder(ctx context.Context, orderID string) error{
	return o.repo.DeleteOrder(ctx, orderID)
}

func mapProtoItemsToOrder(protoItems []*order_proto.OrderItem) []order.OrderItem {
	var items []order.OrderItem
	for _, protoItem := range protoItems {
		items = append(items, order.OrderItem{
			ProductID: protoItem.ProductId,
			Quantity:  protoItem.Quantity,
		})
	}
	return items
}

func mapOrderItemsToProto(items []order.OrderItem) []*order_proto.OrderItem {
	var protoItems []*order_proto.OrderItem
	for _, item := range items {
		protoItems = append(protoItems, &order_proto.OrderItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	return protoItems
}

func mapUnavailableProducts(products []order.UnavailableProduct) []*order_proto.UnavailableProduct {
	var protoProducts []*order_proto.UnavailableProduct
	for _, product := range products {
		protoProducts = append(protoProducts, &order_proto.UnavailableProduct{
			ProductId:        product.ProductID,
			RequestedQuantity: product.RequestedQuantity,
			AvailableQuantity: product.AvailableQuantity,
		})
	}
	return protoProducts
}