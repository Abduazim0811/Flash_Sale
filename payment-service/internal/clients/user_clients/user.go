package userclients

import (
	"context"
	"fmt"
	"log"
	"payment_service/internal/config"
	"payment_service/protos/user_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUsersGrpc() user_proto.UserServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.User.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product", err)
	}

	return user_proto.NewUserServiceClient(conn)
}

func GetUserId(userId string) error {
	_, err := DialUsersGrpc().GetByIdUser(context.Background(), &user_proto.GetUserRequest{Id: userId})
	if err != nil {
		log.Println("users not found: ", err)
		return fmt.Errorf("users not found: %v", err)
	}
	return nil
}

func UpdatePrice(userId string, price float32) error{
	fmt.Println(price)
	_, err := DialUsersGrpc().UpdateUser(context.Background(), &user_proto.UpdateUserReq{Id: userId, Price: price})
	if err != nil {
		log.Println("user price update error: ", err)
		return fmt.Errorf("users price upadte error: %v", err)
	}
	return nil
}
