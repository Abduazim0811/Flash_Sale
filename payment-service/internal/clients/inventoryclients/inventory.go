package inventoryclients

import (
	"context"
	"fmt"
	"log"
	"payment_service/internal/config"
	"payment_service/protos/inventory_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialInventoryGrpc()inventory_proto.InventoryServiceClient{
	c := config.Configuration()
	
	conn, err := grpc.NewClient(c.Inventory.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial inventory grpc error:", err)
	}

	return inventory_proto.NewInventoryServiceClient(conn)
}

func UpdateInventory(productId string, quantity int32) error{
	res, err := DialInventoryGrpc().GetbyIdInventory(context.Background(), &inventory_proto.GetbyIdInventoryReq{
		ProductId: productId,
	})
	if err != nil {
		log.Println("product not found error: ", err)
		return fmt.Errorf("product not found inventory service error: %v", err)
		
	}

	totalQuantity := res.Quantity -quantity

	_, err = DialInventoryGrpc().UpdateInventory(context.Background(), &inventory_proto.UpdateInventoryRequest{
		ProductId: productId,
		Quantity: totalQuantity,
	})
	if  err != nil {
		log.Println("update product quantity error:", err)
		return fmt.Errorf("update product quantity error: %v", err)
	}
	return nil

}