package main

import (
	"cells-auth-server/src/config"
	"cells-auth-server/src/redis"
	"cells-auth-server/src/server"
)

func main() {
	config.LoadConfig("./config.yaml")

	go redis.InitRedis()

	server.StartServer()
}
