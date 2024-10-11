package paymentclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/payment_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialPaymentGrpc() payment_proto.PaymentServiceClient{
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Payment.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal("payment clients error:", err)
	}

	return payment_proto.NewPaymentServiceClient(conn)
}