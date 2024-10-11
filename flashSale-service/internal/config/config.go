package config

import "os"

type Config struct {
	FlashSale struct {
		Host string
		Port string
	}
	Product struct {
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
	c.FlashSale.Host = osGetenv("ORDER_HOST", "tcp")
	c.FlashSale.Port = osGetenv("ORDER_PORT", "flashsale_service:9998")

	c.Product.Port = osGetenv("PRODUCT_PORT", "product_service:7777")

	c.Mongo.Url = osGetenv("MONGO_URL", "mongodb://flashsale_mongo:27017")

	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "flashsale11")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
