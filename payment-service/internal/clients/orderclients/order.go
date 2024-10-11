package orderclients

import (
	"context"
	"fmt"
	"log"
	"payment_service/internal/config"
	"payment_service/protos/order_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialOrderGrpc() order_proto.OrderServiceClient{
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Order.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client order: ", err)
	}

	return order_proto.NewOrderServiceClient(conn)
}

func GetOrder(orderId string)(*order_proto.GetOrderResponse,error){
	res, err := DialOrderGrpc().GetOrder(context.Background(), &order_proto.GetOrderRequest{
		Id: orderId,
	})
	if err != nil{
		return nil, fmt.Errorf("order not found: %v", err)
	}
	_, err = DialOrderGrpc().UpdateOrderStatus(context.Background(), &order_proto.UpdateOrderStatusRequest{
		OrderId: orderId,
		Status: "false",
	})
	if err != nil {
		return nil, fmt.Errorf("update status order error: %v", err)
	}
	return res, nil
}