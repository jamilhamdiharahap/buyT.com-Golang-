package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisConnetion() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	message, err := redis.Ping(ctx).Result()
	if err != nil || message != "PONG" {
		log.Println("Redis Not Connected")
	} else {
		fmt.Println("Redis Connected")
	}
	return redis
}
