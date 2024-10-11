package inventoryclients

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/config"
	"order_service/protos/inventory_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialInventorygrpc() inventory_proto.InventoryServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Inventory.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client inventory: ", err)
	}

	return inventory_proto.NewInventoryServiceClient(conn)
}

func ProductsQuantity(ctx context.Context, product_id string) (*inventory_proto.Inventory, error) {
	res, err := DialInventorygrpc().GetbyIdInventory(ctx, &inventory_proto.GetbyIdInventoryReq{
		ProductId: product_id,
	})
	if err != nil {
		log.Println("product not found:", err)
		return nil, fmt.Errorf("product not found: %v", err)
	}
	return res, nil
}
