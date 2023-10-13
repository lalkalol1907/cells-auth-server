package main

import (
	"cells-auth-server/src/Config"
	"cells-auth-server/src/DB"
	"cells-auth-server/src/Redis"
	"cells-auth-server/src/Server"
	"cells-auth-server/src/gRPC"
	"fmt"
)

func main() {
	Config.LoadConfig("./config.yaml")

	err := Redis.InitRedis()
	defer func() {
		err := Redis.CloseRedis()
		if err != nil {
			print(err)
		}
	}()
	if err != nil {
		return
	}

	err = DB.InitDatabase()
	//defer DB.CloseDatabase()
	if err != nil {
		fmt.Print(err)
		return
	}

	go gRPC.InitServer()

	Server.InitServer()
}
