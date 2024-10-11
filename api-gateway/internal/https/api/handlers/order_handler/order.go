package orderhandler

import (
	"api-gateway/internal/protos/order_proto"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {
	ClientOrder order_proto.OrderServiceClient
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Param order body order_proto.CreateOrderRequest true "Order request body"
// @Success 200 {object} order_proto.CreateOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /order/create [post]
func (o *OrdersHandler) Createorder(c *gin.Context) {
	var req order_proto.CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.CreateOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserOrders godoc
// @Summary Get all orders of a user
// @Description Get all orders made by a specific user
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} order_proto.GetOrdersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /order/user/{id} [get]
func (o *OrdersHandler) GetUserOrders(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var req order_proto.GetOrdersRequest
	req.UserId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.GetOrders(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetOrderById godoc
// @Summary Get order by ID
// @Description Get details of a specific order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} order_proto.GetOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /order/{id} [get]
func (o *OrdersHandler) GetbyIdOrder(c *gin.Context) {
	orderId := c.Param("id")
	var req order_proto.GetOrderRequest
	req.Id = orderId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.GetOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order_proto.GetAllOrdersResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /orders [get]
func (o *OrdersHandler) GetAllOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.GetAllOrders(ctx, &order_proto.GetAllOrdersRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order
// @Tags order
// @Accept json
// @Produce json
// @Param order body order_proto.UpdateOrderStatusRequest true "Update order status request"
// @Success 200 {object} order_proto.UpdateOrderStatusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /order/status [put]
func (o *OrdersHandler) UpdateOrderstatus(c *gin.Context) {
	var req order_proto.UpdateOrderStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.UpdateOrderStatus(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} order_proto.CreateOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /order/{id} [delete]
func (o *OrdersHandler) Deleteorder(c *gin.Context) {
	orderId := c.Param("id")
	var req order_proto.GetOrderRequest
	req.Id = orderId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.DeleteOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
