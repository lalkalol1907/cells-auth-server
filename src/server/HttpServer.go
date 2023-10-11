package server

import (
	"cells-auth-server/src/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() error {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()

	InitControllers(apiRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + config.Cfg.Server.Port,
	}

	fmt.Print("HTTP started")
	return server.ListenAndServe()
}
