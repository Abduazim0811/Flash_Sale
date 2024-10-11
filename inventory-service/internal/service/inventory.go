package service

import (
	"fmt"
	"inventory-service/internal/entity/inventory"
	"inventory-service/internal/infrastructura/repository"
	"inventory-service/protos/inventory_proto"
	"log"

	"github.com/google/uuid"
)

type InventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (i *InventoryService) Createinventory(req *inventory_proto.CreateInventoryReq) (*inventory_proto.CreateInventoryRes, error) {
	inventoryId := uuid.New().String()
	err := i.repo.InsertInventoryPostgres(inventory.CreateInventory{
		InventoryID: inventoryId,
		ProductId:   req.ProductId,
		Quantity:    req.Quantity,
	})
	if err != nil {
		log.Println("create inventory error:", err)
		return nil, fmt.Errorf("create inventory error: %v", err)
	}
	return &inventory_proto.CreateInventoryRes{Message: "inventory created"}, nil
}

func (i *InventoryService) GetAllinventories(req *inventory_proto.GetAllInventoriesReq) (*inventory_proto.GetAllInventoriesRes, error) {
	res, err := i.repo.GetAllInventory()
	if err != nil {
		log.Println("get all inventories error:", err)
		return nil, fmt.Errorf("get all inventories error:  %v", err)
	}

	var inventories []*inventory_proto.Inventory
	for _, inv := range *res {
		inventories = append(inventories, &inventory_proto.Inventory{
			ProductId: inv.ProductID,
			Quantity:  inv.Quantity,
			IsActive:  inv.IsActive,
			CreatedAt: inv.CreatedAt,
			UpdatedAt: inv.UpdatedAt,
		})
	}

	return &inventory_proto.GetAllInventoriesRes{
		Inventories: inventories,
	}, nil
}

func (i *InventoryService) GetbyIdinventory(req *inventory_proto.GetbyIdInventoryReq) (*inventory_proto.Inventory, error) {
	res, err := i.repo.GetByIdInventory(inventory.GetbyIdInventoryReq{ProductId: req.ProductId})
	if err != nil {
		log.Println("get by id inventory error:", err)
		return nil, fmt.Errorf("get by id inventory error: %v", err)
	}

	return &inventory_proto.Inventory{
		ProductId: res.ProductID,
		Quantity:  res.Quantity,
		IsActive:  res.IsActive,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (i *InventoryService) Updateinventory(req *inventory_proto.UpdateInventoryRequest) (*inventory_proto.UpdateInventoryResponse, error) {
	err := i.repo.UpdateInventory(inventory.UpdateInventoryRequest{ProductID: req.ProductId, Quantity: req.Quantity})
	if err != nil {
		log.Println("update inventory error:", err)
		return nil, fmt.Errorf("update inventory error: %v", err)
	}

	return &inventory_proto.UpdateInventoryResponse{Message: "updated"}, nil
}
