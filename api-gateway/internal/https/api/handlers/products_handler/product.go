package productshandler

import (
	"api-gateway/internal/protos/product_proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	ClientProduct product_proto.ProductServiceClient
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags product
// @Accept json
// @Produce json
// @Param product body product_proto.CreateProductReq true "Product request body"
// @Success 200 {object} product_proto.CreateProductRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /product [post]
func (p *ProductsHandler) Createproduct(c *gin.Context) {
	var req product_proto.CreateProductReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.CreateProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProductById godoc
// @Summary Get product by ID
// @Description Get details of a specific product by its ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} product_proto.Product
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /product/{id} [get]
func (p *ProductsHandler) GetbyIdProduct(c *gin.Context) {
	productId := c.Param("id")
	var req product_proto.GetProductReq
	req.ProductId = productId
	

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.GetProductbyId(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListProducts godoc
// @Summary Get all products
// @Description Get a list of all available products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} product_proto.ListProductsRes
// @Failure 500 {object} string
// @Security Bearer
// @Router /products [get]
func (p *ProductsHandler) GetAllProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.ListProducts(ctx, &product_proto.ListProductsReq{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProduct godoc
// @Summary Update product details
// @Description Update information of a specific product by its ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body product_proto.Product true "Product update request body"
// @Success 200 {object} product_proto.UpdateProductRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /product/{id} [put]
func (p *ProductsHandler) Updateproduct(c *gin.Context) {
	productId := c.Param("id")
	var req product_proto.Product
	req.ProductId = productId
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.UpdateProducts(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a specific product by its ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} product_proto.UpdateProductRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /product/{id} [delete]
func (p *ProductsHandler) Deleteproduct(c *gin.Context) {
	productId := c.Param("id")
	var req product_proto.GetProductReq
	req.ProductId = productId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.DeleteProducts(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
