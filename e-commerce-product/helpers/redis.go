package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.ClusterClient

func SetupRedis() {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7000",
			"127.0.0.1:7001",
			"127.0.0.1:7002",
			"127.0.0.1:7003",
			"127.0.0.1:7004",
			"127.0.0.1:7005",
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
