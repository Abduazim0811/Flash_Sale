package repository

import "payment_service/internal/entity/payment"

type PaymentRepository interface {
	InsertPaymentPostgres(req *payment.PaymentRequest) (*payment.GetPaymentRequest, error)
	GetbyIDPaymentPostgres(paymentId string)(*payment.Payment, error)
}
