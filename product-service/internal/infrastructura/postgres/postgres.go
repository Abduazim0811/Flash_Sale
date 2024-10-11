package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"product-service/internal/entity/product"
	"product-service/internal/infrastructura/repository"

	"github.com/Masterminds/squirrel"
)

type ProductPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) repository.ProductRepository {
	return &ProductPostgres{db: db}
}

func (p *ProductPostgres) InsertProduct(req product.CreateProductReq) (*product.CreateProductRes, error) {
	sql, args, err := squirrel.
		Insert("products").
		Columns("id, name, description, price, discount_price, stock").
		Values(req.ProductID, req.Name, req.Description, req.Price, req.DiscountPrice, req.StockQuantity).
		Suffix("Returning id").PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for insert:", err)
		return nil, fmt.Errorf("insert error: %v", err)
	}
	var res product.CreateProductRes
	row := p.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ProductID); err != nil {
		log.Println("Error executing SQL insert:", err)
		return nil, fmt.Errorf("sql insert product error: %v", err)
	}

	return &product.CreateProductRes{ProductID: res.ProductID, Message: "product created"}, nil
}

func (p *ProductPostgres) GetProduct(req product.GetProductReq) (*product.Product, error) {
	sql, args, err := squirrel.
		Select("id, name, description, price, discount_price, stock").
		From("products").
		Where(squirrel.And{
			squirrel.Eq{"id": req.ProductID},
			squirrel.Eq{"is_active": true},
		}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for select:", err)
		return nil, fmt.Errorf("select error: %v", err)
	}

	var prod product.Product
	row := p.db.QueryRow(sql, args...)
	if err := row.Scan(&prod.ProductID, &prod.Name, &prod.Description, &prod.Price, &prod.DiscountPrice, &prod.StockQuantity); err != nil {
		log.Println("Error fetching product from database:", err)
		return nil, fmt.Errorf("sql get product error: %v", err)
	}

	return &prod, nil
}

func (p *ProductPostgres) ListProduct() (*[]product.Product, error) {
	sql, args, err := squirrel.
		Select("id, name, description, price, discount_price, stock").
		From("products").
		Where(squirrel.Eq{"is_active": true}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for list:", err)
		return nil, fmt.Errorf("list products error: %v", err)
	}

	rows, err := p.db.Query(sql, args...)
	if err != nil {
		log.Println("Error executing SQL for list:", err)
		return nil, fmt.Errorf("sql list products error: %v", err)
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var prod product.Product
		if err := rows.Scan(&prod.ProductID, &prod.Name, &prod.Description, &prod.Price, &prod.DiscountPrice, &prod.StockQuantity); err != nil {
			log.Println("Error scanning product row:", err)
			return nil, fmt.Errorf("sql scan product error: %v", err)
		}
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error in rows iteration:", err)
		return nil, err
	}

	return &products, nil
}

func (p *ProductPostgres) UpdateProduct(req product.Product) error {
	filter := squirrel.Eq{"id": req.ProductID}
	log.Println("REQ:", req)

	update := squirrel.Update("products")
	updateFieldsExist := false

	if req.Name != "" {
		update = update.Set("name", req.Name)
		updateFieldsExist = true
	}
	if req.Description != "" {
		update = update.Set("description", req.Description)
		updateFieldsExist = true
	}
	if req.Price != 0 {
		update = update.Set("price", req.Price)
		updateFieldsExist = true
	}
	if req.DiscountPrice != 0 {
		update = update.Set("discount_price", req.DiscountPrice)
		updateFieldsExist = true
	}
	if req.StockQuantity != 0 {
		update = update.Set("stock", req.StockQuantity)
		updateFieldsExist = true
	}

	if !updateFieldsExist {
		return fmt.Errorf("hech qanday yangilanish maydoni mavjud emas")
	}

	sql, args, err := update.Where(filter).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("SQL so'rovini yaratishda xatolik: %v", err)
	}

	_, err = p.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("mahsulotni yangilashda xatolik: %v", err)
	}

	return nil
}

func (p *ProductPostgres) DeleteProduct(req product.GetProductReq) error {
	filter := squirrel.Eq{"id": req.ProductID, "is_active": true}

	sql, args, err := squirrel.
		Update("products").
		Set("is_active", false).
		Where(filter).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL for update:", err)
		return fmt.Errorf("update error: %v", err)
	}

	result, err := p.db.Exec(sql, args...)
	if err != nil {
		log.Println("Error executing SQL for update:", err)
		return fmt.Errorf("sql update product error: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	log.Println("Product marked as inactive with ID:", req.ProductID)
	return nil
}
