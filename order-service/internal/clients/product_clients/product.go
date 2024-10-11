package productclients

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/config"
	"order_service/protos/product_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialProductGrpc() product_proto.ProductServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Product.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product", err)
	}

	return product_proto.NewProductServiceClient(conn)
}

func Products(productId string)(*product_proto.Product, error){
	res, err := DialProductGrpc().GetProductbyId(context.Background(), &product_proto.GetProductReq{ProductId: productId})
	if err != nil{
		return nil, fmt.Errorf("product not found: %v", err)
	}

	return res, nil
}
