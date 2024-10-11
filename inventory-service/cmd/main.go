package main

import (
	"inventory-service/internal/config"
	"inventory-service/internal/connection"
	"inventory-service/internal/infrastructura/postgres"
	"inventory-service/internal/service"
	inventoryservice "inventory-service/inventory_service"
	"inventory-service/protos/inventory_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	db := connection.Database()

	repo := postgres.NewInventoryPostgres(db)

	service := service.NewInventoryService(repo)

	inventory_service := inventoryservice.NewInventoryGrpc(*service)
	server := grpc.NewServer()

	inventory_proto.RegisterInventoryServiceServer(server, inventory_service)
	lis, err := net.Listen(c.Inventory.Host, c.Inventory.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.Inventory.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
