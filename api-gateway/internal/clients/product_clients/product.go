package productclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/product_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DilaProductGrpc() product_proto.ProductServiceClient{
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Product.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("product clients error: ", err)
	}

	return product_proto.NewProductServiceClient(conn)
}