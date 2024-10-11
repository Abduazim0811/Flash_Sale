package config

import "os"

type Config struct {
	Notification struct{
		Host  string
		Port  string
	}
	Order struct {
		Port string
	}
	User struct{
		Port string
	}
	
	Kafka struct {
		Port  string
		Topic string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Notification.Host = osGetenv("NOTIFICATION_HOST", "tcp")
	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:8787")

	c.Order.Port = osGetenv("ORDER_PORT", "order_service:8888")

	c.User.Port = osGetenv("USER_PORT", "user_service:7878")


	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "notification11")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
