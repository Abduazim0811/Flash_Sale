package config

import "os"

type Config struct {
	Order struct{
		Host  string
		Port  string
	}
	Product struct {
		Port string
	}
	User struct{
		Port string
	}
	Inventory struct{
		Port string
	}
	Notification struct{
		Port string
	}
	Mongo struct {
		Url string
	}
	Kafka struct {
		Port  string
		Topic string
	}
}

func Configuration() *Config {
	c := &Config{}
	c.Order.Host = osGetenv("ORDER_HOST", "tcp")
	c.Order.Port = osGetenv("ORDER_PORT", "order_service:8888")

	c.Product.Port = osGetenv("PRODUCT_PORT", "product_service:7777")

	c.User.Port = osGetenv("USER_PORT", "user_service:7878")

	c.Inventory.Port = osGetenv("INVENTORY_PORT", "inventory_service:9999")

	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:8787")

	c.Mongo.Url=osGetenv("MONGO_URL","mongodb://order_mongo:27017")

	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "incomeexpenses17")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
