package Server

import (
	"cells-auth-server/src/Server/Handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitControllers(router *mux.Router) {
	router.Handle("/login", http.HandlerFunc(Handlers.LoginHandler)).Methods("POST")
	router.Handle("/register", http.HandlerFunc(Handlers.RegisterHandler)).Methods("POST")
}

// TODO: Добавить валидацию
