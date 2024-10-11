package userclients

import (
	"api-gateway/internal/config"
	"api-gateway/internal/protos/user_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUsersGrpc() user_proto.UserServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.User.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal("user clients error: ", err)
	}
	return user_proto.NewUserServiceClient(conn)
}
