package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.ClusterClient

func SetupRedis() {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"172.38.0.11:6379",
			"172.38.0.12:6379",
			"172.38.0.13:6379",
			"172.38.0.14:6379",
			"172.38.0.15:6379",
			"172.38.0.16:6379",
		},
	})

	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Error connecting to redis: ", err)
		return
	}
	Logger.Info("PING REDIS: " + result)
	RedisClient = rdb
}
