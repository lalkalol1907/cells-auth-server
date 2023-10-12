package Redis

import (
	"cells-auth-server/src/Config"
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Config.Cfg.Redis.Host + ":" + Config.Cfg.Redis.Port,
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	return err
}

func CloseRedis() error {
	return RedisClient.Close()
}
