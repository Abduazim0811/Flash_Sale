package inventoryclients

import (
	"context"
	"fmt"
	"log"
	"product-service/internal/config"
	"product-service/protos/inventory_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialInventoryGrpc() inventory_proto.InventoryServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Inventory.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product: ", err)
	}
	return inventory_proto.NewInventoryServiceClient(conn)
}

func CreateProductInventory(productId string, quantity int32)error{
	_, err := DialInventoryGrpc().CreateInvetory(context.Background(), &inventory_proto.CreateInventoryReq{
		ProductId: productId,
		Quantity: quantity,
	})
	if err != nil{
		log.Println("error: ", err)
		return fmt.Errorf("error: %v", err)
	}

	return nil
}
