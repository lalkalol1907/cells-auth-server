package main

import (
	"cells-auth-server/src/Config"
	"cells-auth-server/src/DB"
	"cells-auth-server/src/HttpServer"
	"cells-auth-server/src/Redis"
	"cells-auth-server/src/gRPCServer"
)

func main() {
	err := Config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = Redis.InitRedis()
	defer func() {
		err := Redis.CloseRedis()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	err = DB.InitDatabase()
	defer DB.CloseDatabase()
	if err != nil {
		panic(err)
	}

	go func() {
		err := gRPCServer.InitServer()
		if err != nil {
			panic(err)
		}
	}()

	err = HttpServer.InitServer()
	if err != nil {
		panic(err)
	}
}
