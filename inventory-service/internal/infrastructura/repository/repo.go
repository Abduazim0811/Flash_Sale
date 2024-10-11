package repository

import "inventory-service/internal/entity/inventory"

type InventoryRepository interface {
	InsertInventoryPostgres(req inventory.CreateInventory) error
	GetAllInventory() (*[]inventory.Inventory, error)
	GetByIdInventory(req inventory.GetbyIdInventoryReq) (*inventory.Inventory, error)
	UpdateInventory(req inventory.UpdateInventoryRequest) error
}
