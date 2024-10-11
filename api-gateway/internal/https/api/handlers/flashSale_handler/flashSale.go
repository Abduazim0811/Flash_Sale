package flashsalehandler

import (
	"api-gateway/internal/protos/flashSale_proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FlashSalesHandler struct {
	ClientFlashSale flashSale_proto.FlashSaleServiceClient
}

// CreateFlashSale godoc
// @Summary Create a new flash sale
// @Description Create a new flash sale
// @Tags flashsale
// @Accept json
// @Produce json
// @Param flashsale body flashSale_proto.CreateFlashSaleRequest true "Create flash sale request body"
// @Success 200 {object} flashSale_proto.CreateFlashSaleResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale [post]
func (f *FlashSalesHandler) CreateFlashsale(c *gin.Context) {
	var req flashSale_proto.CreateFlashSaleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.CreateFlashSale(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllFlashSales godoc
// @Summary Get all flash sales
// @Description Get a list of all flash sales
// @Tags flashsale
// @Produce json
// @Success 200 {object} flashSale_proto.ListFlashSalesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale [get]
func (f *FlashSalesHandler) GetAllFlashSales(c *gin.Context){
	var req flashSale_proto.ListFlashSalesRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.ListFlashSales(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetActiveFlashSales godoc
// @Summary Get active flash sales
// @Description Get a list of active flash sales
// @Tags flashsale
// @Produce json
// @Success 200 {object} flashSale_proto.GetActiveFlashSalesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale/active [get]
func (f *FlashSalesHandler) GetActiveFlashSales(c *gin.Context){
	var req flashSale_proto.GetActiveFlashSalesRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.GetActiveFlashSales(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateFlashSale godoc
// @Summary Update a flash sale
// @Description Update an existing flash sale
// @Tags flashsale
// @Accept json
// @Produce json
// @Param flashsale body flashSale_proto.UpdateFlashSaleRequest true "Update flash sale request body"
// @Success 200 {object} flashSale_proto.UpdateFlashSaleResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale [put]
func (f *FlashSalesHandler) UpdateFlashsale(c *gin.Context){
	var req flashSale_proto.UpdateFlashSaleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.UpdateFlashSale(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteFlashSale godoc
// @Summary Delete a flash sale
// @Description Delete an existing flash sale
// @Tags flashsale
// @Accept json
// @Produce json
// @Param flashsale body flashSale_proto.DeleteFlashSaleRequest true "Delete flash sale request body"
// @Success 200 {object} flashSale_proto.DeleteFlashSaleResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale [delete]
func (f *FlashSalesHandler) DeleteFlashsale(c *gin.Context){
	var req flashSale_proto.DeleteFlashSaleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.DeleteFlashSale(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PurchaseProduct godoc
// @Summary Purchase a product during a flash sale
// @Description Purchase a product that is part of a flash sale
// @Tags flashsale
// @Accept json
// @Produce json
// @Param purchase body flashSale_proto.PurchaseProductRequest true "Purchase product request body"
// @Success 200 {object} flashSale_proto.PurchaseProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /flashsale/purchase [post]
func (f *FlashSalesHandler) Purchaseproduct(c *gin.Context){
	var req flashSale_proto.PurchaseProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := f.ClientFlashSale.PurchaseProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}