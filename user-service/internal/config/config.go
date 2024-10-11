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
	User struct {
		Host string
		Port string
	}
	Redis struct {
		Host string
		Port string
	}
	Notification struct{
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
	c.Database.Host = osGetenv("DB_HOST", "user_postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "user_service")

	c.User.Host = osGetenv("PRODUCT_HOST", "tcp")
	c.User.Port = osGetenv("PRODUCT_PORT", "user_service:7878")

	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:8787")

	c.Redis.Host = osGetenv("REDIS_HOST", "redis")
	c.Redis.Port = osGetenv("REDIS_PORT", "6379")

	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "user11")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
