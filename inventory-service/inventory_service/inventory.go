package inventoryservice

import (
	"context"
	"fmt"
	"inventory-service/internal/service"
	"inventory-service/protos/inventory_proto"
	"log"
)

type InventoryGrpc struct {
	inventory_proto.UnimplementedInventoryServiceServer
	s service.InventoryService
}

func NewInventoryGrpc(s service.InventoryService) *InventoryGrpc {
	return &InventoryGrpc{s: s}
}

func (i *InventoryGrpc) CreateInvetory(ctx context.Context, req *inventory_proto.CreateInventoryReq) (*inventory_proto.CreateInventoryRes, error){
	res, err:= i.s.Createinventory(req)
	if err != nil {
		log.Println("create inventory error:", err)
		return nil, fmt.Errorf("create inventory error: %v", err)
	}

	return res, nil
}

func (i *InventoryGrpc) GetAllInventories(ctx context.Context,req *inventory_proto.GetAllInventoriesReq) (*inventory_proto.GetAllInventoriesRes, error){
	res, err := i.s.GetAllinventories(req)
	if err != nil{
		log.Println("get all inventories error:", err)
		return nil, fmt.Errorf("get all inventories error: %v", err)
	}

	return res, nil
}

func (i *InventoryGrpc) GetbyIdInventory(ctx context.Context, req *inventory_proto.GetbyIdInventoryReq) (*inventory_proto.Inventory, error){
	res, err := i.s.GetbyIdinventory(req)
	if err != nil {
		log.Println("get by id inventory error:", err)
		return nil, fmt.Errorf("get by id inventory error: %v", err)
	}

	return res, nil
}

func (i *InventoryGrpc) UpdateInventory(ctx context.Context,req *inventory_proto.UpdateInventoryRequest) (*inventory_proto.UpdateInventoryResponse, error){
	res, err := i.s.Updateinventory(req)
	if err != nil {
		log.Println("update inventory error:", err)
		return nil, fmt.Errorf("update inventory error: %v", err)
	}

	return res, nil
}