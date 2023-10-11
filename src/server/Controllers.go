package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitControllers(router *mux.Router) {
	router.Handle("/login", http.HandlerFunc(LoginHandler)).Methods("POST")
	router.Handle("/register", http.HandlerFunc(RegisterHandler)).Methods("POST")
}
