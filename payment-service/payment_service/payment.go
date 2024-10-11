package paymentservice

import (
	"context"
	"fmt"
	"log"
	"payment_service/internal/service"
	"payment_service/protos/payment_proto"
)

type PaymentGrpc struct {
	payment_proto.UnimplementedPaymentServiceServer
	s service.PaymentService
}

func NewPaymentGrpc(s service.PaymentService) *PaymentGrpc {
	return &PaymentGrpc{s: s}
}

func (p *PaymentGrpc) ProcessPayment(ctx context.Context, req *payment_proto.PaymentRequest) (*payment_proto.PaymentResponse, error){
	res, err := p.s.CreatePayment(req)
	if err != nil {
		log.Println("process payment error: ", err)
		return nil, fmt.Errorf("process payment error: %v", err)
	}

	return res, nil
}

func (p *PaymentGrpc) GetPayment(ctx context.Context,req *payment_proto.GetPaymentRequest) (*payment_proto.Payment, error){
	res, err := p.s.GetbyIdpayment(req)
	if err != nil {
		log.Println("get by id payment error: ",err)
		return nil, fmt.Errorf("get by id payment error: %v", err)
	}

	return res, nil
}