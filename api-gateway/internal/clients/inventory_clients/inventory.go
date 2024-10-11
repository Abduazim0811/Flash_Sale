package inventoryclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/inventory_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialInventoryGrpc() inventory_proto.InventoryServiceClient{
	c := config.Configuration()
	conn, err:= grpc.NewClient(c.Inventory.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("inventory clients error:", err)
	}

	return inventory_proto.NewInventoryServiceClient(conn)
}