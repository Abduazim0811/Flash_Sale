package main

import (
	"log"
	"net"
	"order_service/internal/config"
	"order_service/internal/connections"
	"order_service/internal/infrastructura/mongodb"
	"order_service/internal/service"
	orderservice "order_service/order_service"
	"order_service/protos/order_proto"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	client, collection, err := connections.NewMongodb()
	if err != nil {
		log.Fatal("mongo connection error:", err)
	}

	repo := mongodb.NewOrderMongodb(client, collection)
	service := service.NewOrderService(repo)
	order_service := orderservice.NewOrderGrpc(*service)

	server := grpc.NewServer()
	order_proto.RegisterOrderServiceServer(server, order_service)
	lis, err := net.Listen(c.Order.Host, c.Order.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.Order.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
