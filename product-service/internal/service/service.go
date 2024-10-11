package service

import (
	"fmt"
	"log"
	inventoryclients "product-service/internal/clients/inventory_clients"
	"product-service/internal/entity/product"
	"product-service/internal/infrastructura/redis"
	"product-service/internal/infrastructura/repository"
	"product-service/protos/product_proto"
	"time"

	"github.com/google/uuid"
)

type ProductService struct {
	repo  repository.ProductRepository
	cache *redis.RedisClient
}

func NewProductService(repo repository.ProductRepository, cache *redis.RedisClient) *ProductService {
	return &ProductService{repo: repo, cache: cache}
}
func (p *ProductService) Createproduct(req product.CreateProductReq) (*product.CreateProductRes, error) {
	productID := uuid.New().String()
	req.ProductID = productID

	res, err := p.repo.InsertProduct(req)
	if err != nil {
		return nil, err
	}

	err = p.cache.Set(res.ProductID, req, 0)
	if err != nil {
		log.Println("Failed to cache product:", err)
	}
	err = inventoryclients.CreateProductInventory(req.ProductID, req.StockQuantity)
	if err != nil {
		log.Println("create inventory error: ", err)
		return nil, fmt.Errorf("create inventory error: %v", err)
	}

	return res, nil
}

func (p *ProductService) Getproduct(req product.GetProductReq) (*product.Product, error) {
	var cachedProduct product.Product

	err := p.cache.Get(req.ProductID, &cachedProduct)
	if err == nil && cachedProduct.ProductID != "" {
		fmt.Println("Keshdan olingan mahsulot:", cachedProduct.ProductID, cachedProduct.Name, cachedProduct.Price)
		return &cachedProduct, nil
	} else if err != nil {
		log.Printf("Redis'dan mahsulotni olishda xato: %v", err)
	}

	res, err := p.repo.GetProduct(req)
	if err != nil {
		return nil, err
	}

	err = p.cache.Set(res.ProductID, res, 10*time.Minute)
	if err != nil {
		log.Println("Mahsulotni Redis keshiga qo'shishda xato:", err)
	}

	return res, nil
}

func (p *ProductService) Listproducts() (*[]product.Product, error) {
	return p.repo.ListProduct()
}

func (p *ProductService) Updateproduct(req product.Product) error {
	err := p.repo.UpdateProduct(req)
	if err != nil {
		return err
	}

	err = p.cache.Delete(req.ProductID)
	if err != nil {
		log.Println("Failed to remove product from cache:", err)
	}

	return nil
}

func (p *ProductService) Deleteproduct(req product.GetProductReq) error {
	err := p.repo.DeleteProduct(req)
	if err != nil {
		return err
	}

	err = p.cache.Delete(req.ProductID)
	if err != nil {
		log.Println("Failed to remove product from cache:", err)
	}

	return nil
}

func (p *ProductService) GetProductorder(req *product_proto.GetProductOrderReq) (*product_proto.Product, error) {
	res, err := p.repo.GetProduct(product.GetProductReq{ProductID: req.ProductId})
	if err != nil {
		return nil, fmt.Errorf("get product order error: %v", err)
	}
	return &product_proto.Product{
		ProductId:     res.ProductID,
		Name:          res.Name,
		Description:   res.Description,
		Price:         res.Price,
		DiscountPrice: res.DiscountPrice,
		StockQuantity: res.StockQuantity,
	}, nil
}
