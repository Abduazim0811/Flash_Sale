package service

import (
	"fmt"
	"log"
	userclients "payment_service/internal/clients/user_clients"
	"payment_service/internal/entity/payment"
	"payment_service/internal/infrastructura/repository"
	"payment_service/protos/payment_proto"
)

type PaymentService struct{
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *PaymentService{
	return &PaymentService{repo: repo}
}

func (p *PaymentService) CreatePayment(req *payment_proto.PaymentRequest)(*payment_proto.PaymentResponse, error){
	err := userclients.GetUserId(req.UserId)
	if err != nil {
		log.Println("users not found:", err)
		return nil, fmt.Errorf("user not found: %v", err)
	}
	res, err := p.repo.InsertPaymentPostgres(&payment.PaymentRequest{
		OrderID: req.OrderId,
		UserID: req.UserId,
		Amount: req.Amount,
	})
	if err != nil {
		log.Println("create payment error: ", err)
		return nil, fmt.Errorf("create payment error: %v", err)
	}

	return &payment_proto.PaymentResponse{PaymentId: res.PaymentID, Message: "created"}, nil
}

func (p *PaymentService) GetbyIdpayment(req *payment_proto.GetPaymentRequest)(*payment_proto.Payment, error){
	res, err := p.repo.GetbyIDPaymentPostgres(req.PaymentId)
	if err != nil {
		log.Println("get by id payment error: ", err)
		return nil, fmt.Errorf("get by id payment error: %v", err)
	}

	return &payment_proto.Payment{
		Id: res.ID,
		OrderId: res.OrderID,
		UserId: res.UserID,
		Amount: res.Amount,
		Status: res.Status,
		CreatedAt: res.CreatedAt,
	}, nil
}