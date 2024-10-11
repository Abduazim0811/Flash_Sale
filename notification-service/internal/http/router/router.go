package router

import (
	"fmt"
	"log"
	"net"
	"notification-service/internal/config"
	"notification-service/internal/http/websocket"
	"notification-service/internal/infrastructura/kafka"
	"notification-service/internal/service"
	"notification-service/protos/notification_proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Router() {
	c := config.Configuration()

	kafkaProducer, err := kafka.NewProducer([]string{c.Kafka.Port})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	go func() {
		Grpc()
	}()

	wsManager, err := websocket.NewFlashSaleWebSocket([]string{c.Kafka.Port})
	if err != nil {
		log.Fatalf("Failed to create FlashSale WebSocket: %v", err)
	}

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		wsManager.HandleWebSocket(c)
	})

	log.Println("WebSocket server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start WebSocket server: %v", err)
	}
}

func Grpc() {
	c := config.Configuration()

	kafkaProducer, err := kafka.NewProducer([]string{c.Kafka.Port})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	listener, err := net.Listen(c.Notification.Host, c.Notification.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	notificationServer := service.NewNotificationServer(kafkaProducer)
	notification_proto.RegisterNotificationServiceServer(s, notificationServer)
	reflection.Register(s)

	fmt.Printf("gRPC server started on %s\n", c.Notification.Port)

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
