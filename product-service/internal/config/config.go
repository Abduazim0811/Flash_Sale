package config

import "os"

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBname   string
	}
	Product struct {
		Host string
		Port string
	}
	Inventory struct{
		Port string
	}
	Redis struct {
		Host string
		Port string
	}
	Kafka struct {
		Port  string
		Topic string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Database.User = osGetenv("DB_USER", "postgres")
	c.Database.Password = osGetenv("DB_PASSWORD", "Abdu0811")
	c.Database.Host = osGetenv("DB_HOST", "product_postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "product_service")

	c.Product.Host = osGetenv("PRODUCT_HOST", "tcp")
	c.Product.Port = osGetenv("PRODUCT_PORT", "product_service:7777")

	c.Inventory.Port = osGetenv("INVENTORY_PORT", "inventory_service:9999")

	c.Redis.Host = osGetenv("REDIS_HOST", "redis")
	c.Redis.Port = osGetenv("REDIS_PORT", "6379")

	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "product11")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
