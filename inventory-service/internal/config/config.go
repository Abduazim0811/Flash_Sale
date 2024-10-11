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
	Inventory struct {
		Host string
		Port string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Database.User = osGetenv("DB_USER", "postgres")
	c.Database.Password = osGetenv("DB_PASSWORD", "Abdu0811")
	c.Database.Host = osGetenv("DB_HOST", "inventory_postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "product_service")

	c.Inventory.Host = osGetenv("INVENTORY_HOST", "tcp")
	c.Inventory.Port = osGetenv("INVENTORY_PORT", "inventory_service:9999")
	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
