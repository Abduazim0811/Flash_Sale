package main

import (
	"log"
	"net"
	"product-service/internal/config"
	"product-service/internal/connections"
	"product-service/internal/infrastructura/kafka"
	"product-service/internal/infrastructura/postgres"
	"product-service/internal/infrastructura/redis"
	"product-service/internal/service"
	productservice "product-service/product-service"
	"product-service/protos/product_proto"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	db := connections.Database()

	r := redis.NewRedisClient()

	repo := postgres.NewProductPostgres(db)
	service := service.NewProductService(repo, r)

	product_service := productservice.NewProductGRPC(*service)
	con := kafka.NewConsumer(service)
	go func(){
		con.Consumer()
	}()
	server := grpc.NewServer()
	product_proto.RegisterProductServiceServer(server, product_service)
	lis, err := net.Listen(c.Product.Host, c.Product.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.Product.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
