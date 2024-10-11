package flashsaleclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/flashSale_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialFlashSaleGrpc() flashSale_proto.FlashSaleServiceClient{
	c := config.Configuration()
	conn, err := grpc.NewClient(c.FlashSale.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("flash sale clients error:", err)
	}

	return flashSale_proto.NewFlashSaleServiceClient(conn)
}