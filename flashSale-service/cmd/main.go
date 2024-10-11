package main

import (
	flashsaleservice "flashsale-service/flashSale_service"
	"flashsale-service/internal/config"
	"flashsale-service/internal/connections"
	"flashsale-service/internal/infrastructura/mongodb"
	"flashsale-service/internal/service"
	"flashsale-service/protos/flashSale_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	client, collection, err := connections.NewMongodb()
	if err != nil {
		log.Fatal("mongodb error: ", err)
	}
	repo := mongodb.NewFlashSaleMongodb(client, collection)
	servc := service.NewFlashSaleService(repo)
	flashsale_service := flashsaleservice.NewFlashSaleGrpc(*servc)
	server := grpc.NewServer()
	flashSale_proto.RegisterFlashSaleServiceServer(server, flashsale_service)
	lis, err := net.Listen(c.FlashSale.Host, c.FlashSale.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.FlashSale.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
