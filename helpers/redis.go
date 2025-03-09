package helpers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_HOST", "localhost:6379"),
		DB:  0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Failed to connect to Redis: ", err)
		return
	}
	Logger.Info("PING REDIS: " + ping)

	RedisClient = client
}