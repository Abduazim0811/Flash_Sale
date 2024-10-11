package service

import (
	"context"
	"fmt"
	"notification-service/internal/infrastructura/kafka"
	"notification-service/protos/notification_proto"
)

type NotificationServer struct {
	notification_proto.UnimplementedNotificationServiceServer
	kafkaProducer *kafka.Producer
}

func NewNotificationServer(producer *kafka.Producer) *NotificationServer {
	return &NotificationServer{kafkaProducer: producer}
}

func (s *NotificationServer) SendNotification(ctx context.Context, req *notification_proto.SendNotificationRequest) (*notification_proto.SendNotificationResponse, error) {
	err := s.kafkaProducer.PublishNotification(req.UserId, req.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %v", err)
	}
	return &notification_proto.SendNotificationResponse{Success: true}, nil
}

func (s *NotificationServer) SubscribeToNotifications(req *notification_proto.SubscribeRequest, stream notification_proto.NotificationService_SubscribeToNotificationsServer) error {
	messages := s.kafkaProducer.Subscribe(req.UserId)

	for msg := range messages {
		err := stream.Send(&notification_proto.NotificationMessage{
			Message:   msg.Message,
		})
		if err != nil {
			return fmt.Errorf("failed to send notification: %v", err)
		}
	}
	return nil
}
