package repository

import "product-service/internal/entity/product"

type ProductRepository interface{
	InsertProduct(req product.CreateProductReq)(*product.CreateProductRes, error)
	GetProduct(req product.GetProductReq)(*product.Product, error)
	ListProduct()(*[]product.Product, error)
	UpdateProduct(req product.Product)(error)
	DeleteProduct(req product.GetProductReq)(error)
}