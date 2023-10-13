package Server

import (
	"cells-auth-server/src/Config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitServer() error {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()

	InitControllers(apiRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + Config.Cfg.HttpServer.Port,
	}

	fmt.Print("HTTP started")
	return server.ListenAndServe()
}
