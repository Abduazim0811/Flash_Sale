package orderclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/order_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialOrderGrpc()order_proto.OrderServiceClient{
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Order.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("order clients error:", err)
	}

	return order_proto.NewOrderServiceClient(conn)
}