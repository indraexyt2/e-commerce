package helpers

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func SetupRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_HOST"),
		DB:   0,
	})

	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Error connecting to redis: ", err)
		return
	}
	RedisClient = rdb
	Logger.Info("PING REDIS: " + result)
	RedisClient = rdb
}
