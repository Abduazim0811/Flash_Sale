package notificationclients

import (
	"log"
	"payment_service/internal/config"
	"payment_service/protos/notification_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialNotificationGrpc() notification_proto.NotificationServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Notification.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc notification error:", err)
	}
	log.Println("ulandi")
	return notification_proto.NewNotificationServiceClient(conn)
}
