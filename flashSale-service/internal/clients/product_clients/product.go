package productclients

import (
	"context"
	"flashsale-service/internal/config"
	"flashsale-service/protos/product_proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialProductgrpc() product_proto.ProductServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Product.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product: ", err)
	}

	return product_proto.NewProductServiceClient(conn)
}

func GetProductbyId(id string)(*product_proto.Product, error){
	res, err:= DialProductgrpc().GetProductbyId(context.Background(), &product_proto.GetProductReq{ProductId: id})
	if err != nil {
		log.Println("product not found: ", err)
		return nil, fmt.Errorf("product not found: %v", err)
	}

	return res, nil
}

func UpdateProductDiscountPrice(productid string, discountprice float32)error{
	_, err := DialProductgrpc().UpdateProducts(context.Background(), &product_proto.Product{ProductId: productid,DiscountPrice: discountprice})
	if err != nil {
		log.Println("update product error: ", err)
		return fmt.Errorf("update product error: %v", err)
	}
	return nil
}
