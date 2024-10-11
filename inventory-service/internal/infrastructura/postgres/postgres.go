package postgres

import (
	"database/sql"
	"fmt"
	"inventory-service/internal/entity/inventory"
	"inventory-service/internal/infrastructura/repository"
	"log"

	"github.com/Masterminds/squirrel"
)

type InventoryPostgers struct {
	db  *sql.DB
}

func NewInventoryPostgres(db *sql.DB) repository.InventoryRepository {
	return &InventoryPostgers{db: db}
}

func (i *InventoryPostgers) InsertInventoryPostgres(req inventory.CreateInventory) error{
	sql, args, err := squirrel.
		Insert("inventory").
		Columns("id, product_id, quantity").
		Values(req.InventoryID, req.ProductId, req.Quantity).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for insert:", err)
		return fmt.Errorf("error generating SQL for insert: %v", err)
	}

	_, err = i.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing SQL insert:", err)
		return fmt.Errorf("error executing SQL insert: %v", err)
	}
	return nil
}

func (i *InventoryPostgers) GetAllInventory() (*[]inventory.Inventory, error) {
	sql, args, err := squirrel.
		Select("product_id, quantity, created_at, updated_at").
		From("inventory").
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Println("Error generating SQL for get all inventories:", err)
		return nil, fmt.Errorf("error generating SQL for get all inventories: %v", err)
	}

	rows, err := i.db.Query(sql, args...)
	if err != nil {
		log.Println("Error executing SQL for get all inventories:", err)
		return nil, fmt.Errorf("error executing SQL for get all inventories: %v", err)
	}
	defer rows.Close()

	var inventories []inventory.Inventory
	for rows.Next() {
		var inv inventory.Inventory
		if err := rows.Scan(&inv.ProductID, &inv.Quantity, &inv.CreatedAt, &inv.UpdatedAt); err != nil {
			log.Println("Error scanning inventory row:", err)
			return nil, fmt.Errorf("error scanning inventory row: %v", err)
		}
		inventories = append(inventories, inv)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in rows iteration:", err)
		return nil, err
	}

	return &inventories, nil
}

func (i *InventoryPostgers) GetByIdInventory(req inventory.GetbyIdInventoryReq) (*inventory.Inventory, error) {
	sql, args, err := squirrel.
		Select("product_id, quantity, created_at, updated_at").
		From("inventory").
		Where(squirrel.Eq{"product_id": req.ProductId}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Println("Error generating SQL for get inventory by id:", err)
		return nil, fmt.Errorf("error generating SQL for get inventory by id: %v", err)
	}

	row := i.db.QueryRow(sql, args...)

	var inv inventory.Inventory
	if err := row.Scan(&inv.ProductID, &inv.Quantity, &inv.CreatedAt, &inv.UpdatedAt); err != nil {
		log.Println("Error scanning inventory row:", err)
		return nil, fmt.Errorf("error scanning inventory row: %v", err)
	}

	return &inv, nil
}

func (i *InventoryPostgers) UpdateInventory(req inventory.UpdateInventoryRequest) error{
	sql, args, err := squirrel.
		Update("inventory").
		Set("quantity", req.Quantity).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{"product_id": req.ProductID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Println("Error generating SQL for update inventory:", err)
		return  fmt.Errorf("error generating SQL for update inventory: %v", err)
	}

	_, err = i.db.Exec(sql, args...)
	if err != nil {
		log.Println("Error executing SQL update:", err)
		return fmt.Errorf("error executing SQL update: %v", err)
	}

	return nil
}
