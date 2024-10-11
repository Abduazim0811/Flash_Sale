package main

import (
	"log"
	"net"
	"payment_service/internal/config"
	"payment_service/internal/connections"
	"payment_service/internal/infrastructura/postgres"
	"payment_service/internal/service"
	paymentservice "payment_service/payment_service"
	"payment_service/protos/payment_proto"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	db := connections.Database()
	repo := postgres.NewPaymentPostgres(db)

	service := service.NewPaymentService(repo)
	payment_service := paymentservice.NewPaymentGrpc(*service)
	server := grpc.NewServer()
	payment_proto.RegisterPaymentServiceServer(server, payment_service)
	lis, err := net.Listen(c.Payment.Host, c.Payment.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.Payment.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
