package paymentshandler

import (
	"api-gateway/internal/protos/payment_proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentsHandler struct {
	ClientPayment payment_proto.PaymentServiceClient
}

// ProcessPayment godoc
// @Summary Process a payment
// @Description Process a new payment
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body payment_proto.PaymentRequest true "Payment request body"
// @Success 200 {object} payment_proto.PaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /payment/process [post]
func (p *PaymentsHandler) ProcessPayments(c *gin.Context) {
	var req payment_proto.PaymentRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientPayment.ProcessPayment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPayment godoc
// @Summary Get payment by ID
// @Description Get details of a specific payment by its ID
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} payment_proto.Payment
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /payment/{id} [get]
func (p *PaymentsHandler) Getpayment(c *gin.Context) {
	paymentId := c.Param("id")
	var req payment_proto.GetPaymentRequest
	req.PaymentId = paymentId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientPayment.GetPayment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
