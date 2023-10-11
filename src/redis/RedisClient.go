package redis

import (
	"cells-auth-server/src/config"
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Host + ":" + config.Cfg.Redis.Port,
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	return err
}
