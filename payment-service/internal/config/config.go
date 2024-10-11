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
	Payment struct{
		Host  string
		Port  string
	}
	Order struct {
		Port string
	}
	User struct{
		Port string
	}
	Notification struct{
		Port string
	}
	Inventory struct{
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
	c.Database.Host = osGetenv("DB_HOST", "payment_postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "payment_service")

	c.Payment.Host = osGetenv("PAYMENT_HOST", "tcp")
	c.Payment.Port = osGetenv("PAYMENT_PORT", "payment_service:7778")

	c.Order.Port = osGetenv("ORDER_PORT", "order_service:8888")

	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:8787")

	c.Inventory.Port = osGetenv("INVENTORY_PORT", "inventory_service:9999")

	c.User.Port = osGetenv("USER_PORT", "user_service:7878")


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
