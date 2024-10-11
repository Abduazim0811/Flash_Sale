package kafka

import (
	"context"
	"fmt"
	"log"
	"notification-service/internal/entity/notification"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(brokers []string) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ProducerBatchMaxBytes(1000), // Producer uchun max batch hajmi
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %v", err)
	}

	return &Producer{client: client}, nil
}

func (p *Producer) PublishNotification(userID, message string) error {
    log.Printf("Publishing notification for user: %s, message: %s", userID, message)  // Log qo'shish
    record := &kgo.Record{
        Topic: "notification-topic",
        Key:   []byte(userID),
        Value: []byte(message),
    }

    ctx := context.Background()
    p.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
        if err != nil {
            log.Printf("failed to produce record: %v", err)
        }else{
			log.Printf("Produced message for userID: %s, message: %s", userID, message)
		}
    })

    return nil
}


func (p *Producer) Subscribe(userID string) chan *notification.NotificationMessage {
	messages := make(chan *notification.NotificationMessage)

	go func() {
		for {
			fetches := p.client.PollFetches(context.Background())
			if fetches.Err() != nil {
				log.Printf("failed to fetch records: %v", fetches.Err())
				close(messages)
				break
			}

			fetches.EachRecord(func(record *kgo.Record) {
				messages <- &notification.NotificationMessage{
					Message:   string(record.Value),
				}
			})
		}
	}()

	return messages
}

func (p *Producer) Close() {
	p.client.Close()
}
