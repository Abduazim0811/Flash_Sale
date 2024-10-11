package config

import "os"

type Config struct {
	ApiGateway struct{
		Port string
	}
	User struct {
		Port string
	}
	Product struct {
		Port string
	}
	Inventory struct {
		Port string
	}
	Order struct {
		Port string
	}
	Payment struct {
		Port string
	}
	FlashSale struct {
		Port string
	}
	Notification struct {
		Port string
	}
	Kafka struct {
		Port  string
		Topic string
	}
}

func Configuration() *Config {
	c := &Config{}
	c.ApiGateway.Port = osGetenv("API_GATEWAY", "api_gateway:9876")
	
	c.User.Port = osGetenv("USER_PORT", "user_service:7878")
	
	c.Order.Port = osGetenv("ORDER_PORT", "order_service:8888")

	c.Product.Port = osGetenv("PRODUCT_PORT", "product_service:7777")

	c.Inventory.Port = osGetenv("INVENTORY_PORT", "inventory_service:9999")

	c.FlashSale.Port = osGetenv("FLASHSALE_PORT", "flashsale_service:9998")

	c.Payment.Port = osGetenv("PAYMENT_PORT", "payment_service:7778")

	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:8787")
	
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
