package inventoryhandler

import (
	"api-gateway/internal/protos/inventory_proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	ClientInventory inventory_proto.InventoryServiceClient
}

// CreateInventory godoc
// @Summary Create a new inventory
// @Description Create a new inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body inventory_proto.CreateInventoryReq true "inventory request body"
// @Success 200 {object} inventory_proto.CreateInventoryRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /inventory [post]
func (i *InventoryHandler) Createinventory(c *gin.Context) {
	var req inventory_proto.CreateInventoryReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := i.ClientInventory.CreateInvetory(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetbyIdinventory godoc
// @Summary Get inventory by ID
// @Description Get an inventory by its ID
// @Tags inventory
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} inventory_proto.Inventory
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /inventory/{id} [get]
func (i *InventoryHandler) GetbyIdinventory(c *gin.Context) {
	productId := c.Param("id")
	var req inventory_proto.GetbyIdInventoryReq
	req.ProductId = productId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := i.ClientInventory.GetbyIdInventory(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllInventory godoc
// @Summary Get all inventories
// @Description Get all inventories
// @Tags inventory
// @Produce json
// @Success 200 {object} inventory_proto.GetAllInventoriesRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /inventory [get]
func (i *InventoryHandler) GetAllinventory(c *gin.Context) {
	var req inventory_proto.GetAllInventoriesReq

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := i.ClientInventory.GetAllInventories(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateInventory godoc
// @Summary Update an inventory
// @Description Update an existing inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body inventory_proto.UpdateInventoryRequest true "Update inventory request body"
// @Success 200 {object} inventory_proto.UpdateInventoryResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /inventory [put]
func (i *InventoryHandler) Updateinventory(c *gin.Context) {
	var req inventory_proto.UpdateInventoryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := i.ClientInventory.UpdateInventory(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
