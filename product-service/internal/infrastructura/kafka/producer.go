package kafka

import (
	"context"
	"product-service/internal/config"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProducer struct {
	Client *kgo.Client
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{Client: client}, nil
}


func (p *KafkaProducer) SendMessage( key, value string) error {
	c := config.Configuration()
	record := &kgo.Record{
		Topic: c.Kafka.Topic,
		Key:   []byte(key),
		Value: []byte(value),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return p.Client.ProduceSync(ctx, record).FirstErr()
}

func (p *KafkaProducer) Close() {
	p.Client.Close()
}
