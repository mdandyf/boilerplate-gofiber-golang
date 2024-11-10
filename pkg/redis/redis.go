package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func Client(addr, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Ping the Redis server to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisClient{
		client: client,
	}, nil
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %v", err)
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key does not exist in Redis")
		}
		return "", fmt.Errorf("failed to get value from Redis: %v", err)
	}
	return val, nil
}

func (r *RedisClient) Delete(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete value from Redis: %v", err)
	}
	return nil
}

func (r *RedisClient) Close() error {
	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("failed to close Redis client: %v", err)
	}
	return nil
}
