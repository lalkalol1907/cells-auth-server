package main

import (
	"cells-auth-server/src/Config"
	"cells-auth-server/src/Redis"
	"cells-auth-server/src/Server"
)

func main() {
	Config.LoadConfig("./config.yaml")

	err := Redis.InitRedis()

	if err != nil {
		return
	}

	Server.StartServer()
}
