package productservice

import (
	"context"
	"fmt"
	"log"
	"product-service/internal/entity/product"
	"product-service/internal/service"
	"product-service/protos/product_proto"
)

type ProductGRPC struct {
	product_proto.UnimplementedProductServiceServer
	s service.ProductService
}

func NewProductGRPC(s service.ProductService) *ProductGRPC {
	return &ProductGRPC{s: s}
}

func (p *ProductGRPC) CreateProduct(ctx context.Context, req *product_proto.CreateProductReq) (*product_proto.CreateProductRes, error){
	res, err := p.s.Createproduct(product.CreateProductReq{
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		DiscountPrice: req.DiscountPrice,
		StockQuantity: req.StockQuantity,
	})
	if err != nil {
		log.Println("product create error: ", err)
		return nil, fmt.Errorf("product create error: %v", err)
	}

	return &product_proto.CreateProductRes{ProductId: res.ProductID, Message: res.Message}, nil
}

func (p *ProductGRPC) GetProductbyId(ctx context.Context,req *product_proto.GetProductReq) (*product_proto.Product, error){
	res, err := p.s.Getproduct(product.GetProductReq{ProductID: req.ProductId})
	if err != nil {
		log.Println("get product error: ", err)
		return nil, fmt.Errorf("get product error: %v", err)
	}
	return &product_proto.Product{
		ProductId: res.ProductID,
		Name: res.Name,
		Description: res.Description,
		Price: res.Price,
		DiscountPrice: res.DiscountPrice,
		StockQuantity: res.StockQuantity,
	}, nil
}

func (p *ProductGRPC) ListProducts(ctx context.Context,req *product_proto.ListProductsReq) (*product_proto.ListProductsRes, error){
	products, err := p.s.Listproducts()
	if err != nil {
		log.Println("list products error: ", err)
		return nil, fmt.Errorf("list products error: %v", err)
	}

	var productProtoList []*product_proto.Product
	for _, prod := range *products {
		productProtoList = append(productProtoList, &product_proto.Product{
			ProductId:    prod.ProductID,
			Name:         prod.Name,
			Description:  prod.Description,
			Price:        prod.Price,
			DiscountPrice: prod.DiscountPrice,
			StockQuantity: prod.StockQuantity,
		})
	}

	return &product_proto.ListProductsRes{Product: productProtoList}, nil
}

func (p *ProductGRPC) UpdateProducts(ctx context.Context,req *product_proto.Product) (*product_proto.UpdateProductRes, error){
	err := p.s.Updateproduct(product.Product{
		ProductID: req.ProductId,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		StockQuantity: req.StockQuantity,
		
	})
	if err != nil {
		log.Println("update product error: ", err)
		return nil, fmt.Errorf("update product error: %v", err)
	}
	return &product_proto.UpdateProductRes{Message: "product updated successfully"}, nil
}

func (p *ProductGRPC) DeleteProducts(ctx context.Context, req *product_proto.GetProductReq) (*product_proto.UpdateProductRes, error){
	err := p.s.Deleteproduct(product.GetProductReq{ProductID: req.ProductId})
	if err != nil {
		log.Println("delete product error: ", err)
		return nil, fmt.Errorf("delete product error: %v", err)
	}
	return &product_proto.UpdateProductRes{Message: "product deleted successfully"}, nil
}

func (p *ProductGRPC) GetProductOrder(ctx context.Context, req *product_proto.GetProductOrderReq) (*product_proto.Product, error){
	res, err := p.s.GetProductorder(req)
	if err != nil {
		return nil, fmt.Errorf("get product order error: %v", err)
	}

	return res, nil
}