package orderservice

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/service"
	"order_service/protos/order_proto"
)

type OrderGrpc struct {
	order_proto.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderGrpc(s service.OrderService) *OrderGrpc {
	return &OrderGrpc{service: s}
}

func (o *OrderGrpc) CreateOrder(ctx context.Context, req *order_proto.CreateOrderRequest) (*order_proto.CreateOrderResponse, error) {
	res,err := o.service.Createorder(ctx, req)
	if err != nil {
		log.Println("create order error:  ", err)
		return nil, fmt.Errorf("create order error: %v", err)
	}

	return &order_proto.CreateOrderResponse{
		Status: res.Status,
		OrderId: res.OrderId,
		Message: res.Message,
		TotalAmount: res.TotalAmount,
		UnavailableProducts: res.UnavailableProducts,
	}, nil
}

func (o *OrderGrpc) GetAllOrders(ctx context.Context,req *order_proto.GetAllOrdersRequest) (*order_proto.GetAllOrdersResponse, error){
	res, err := o.service.GetAllorders(ctx)
	if err != nil {
		log.Println("get all orders error: ", err)
		return nil, fmt.Errorf("get all orders error: %v", err)
	}

	return res,nil
}

func (o *OrderGrpc) GetOrder(ctx context.Context,req *order_proto.GetOrderRequest) (*order_proto.GetOrderResponse, error){
	res, err := o.service.Getorder(ctx, req.Id)
	if err != nil {
		log.Println("get order error:", err)
		return nil, fmt.Errorf("get order error: %v", err)
	}

	return res, nil
}

func (o *OrderGrpc) GetOrders(ctx context.Context,req *order_proto.GetOrdersRequest) (*order_proto.GetOrdersResponse, error){
	res, err := o.service.GetUserorders(ctx, req.UserId)
	if err != nil {
		log.Println("get orders error:", err)
		return nil, fmt.Errorf("get orders  error: %v", err)
	}

	return res, nil
}

func (o *OrderGrpc) UpdateOrderStatus(ctx context.Context,req *order_proto.UpdateOrderStatusRequest) (*order_proto.UpdateOrderStatusResponse, error){
	res, err := o.service.UpdateorderStatus(ctx, req.OrderId, req.Status)
	if err != nil{
		log.Println("update order status error:", err)
		return nil, fmt.Errorf("update order status error: %v", err)
	}

	return &order_proto.UpdateOrderStatusResponse{Order: res}, err
}

func (o *OrderGrpc) DeleteOrder(ctx context.Context,req *order_proto.GetOrderRequest) (*order_proto.CreateOrderResponse, error){
	err := o.service.Deleteorder(ctx, req.Id)
	if err != nil {
		log.Println("delete order error:", err)
		return nil, fmt.Errorf("delete order error: %v", err)
	}

	return &order_proto.CreateOrderResponse{Message: "Order deleted"}, nil
}