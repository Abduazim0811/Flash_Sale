package main

import (
	"fmt"
	"log"
	"net"
	notificationclients "user-service/internal/clients/notificationClients"
	"user-service/internal/config"
	"user-service/internal/connections"
	"user-service/internal/infrastructura/postgres"
	"user-service/internal/infrastructura/redis"
	"user-service/internal/service"
	"user-service/protos/user_proto"
	userservice "user-service/user_service"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()

	db := connections.Database()
	n := notificationclients.DialNotificationGrpc()

	repo := postgres.NewUserPostgres(db,n)
	redisaddr := fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
	r := redis.NewRedisClient(redisaddr, "", 0)
	service := service.NewUserService(repo, r)

	users_service := userservice.NewUsersGrpc(*service)
	server := grpc.NewServer()
	user_proto.RegisterUserServiceServer(server, users_service)
	lis, err := net.Listen(c.User.Host, c.User.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.User.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
