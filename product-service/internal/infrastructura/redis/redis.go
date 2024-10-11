package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"product-service/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	cfg := config.Configuration()
	redisAddr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value (key: %s): %v", key, err)
	}

	err = r.client.Set(context.Background(), key, data, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key '%s' in Redis: %v", key, err)
	}

	return nil
}




func (r *RedisClient) Get(key string, dest interface{}) error {
	if r.client == nil {
		return fmt.Errorf("redis mijoz mavjud emas")
	}

	data, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("kalit '%s' topilmadi", key)
		}
		return fmt.Errorf("redis'dan '%s' kaliti uchun qiymatni olishda xato: %v", key, err)
	}

	err = json.Unmarshal([]byte(data), dest)
	if err != nil {
		return fmt.Errorf("qiymatni deserializatsiya qilishda xato (key: %s): %v", key, err)
	}

	return nil
}


func (r *RedisClient) Delete(key string) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}

	return r.client.Del(context.Background(), key).Err()
}
