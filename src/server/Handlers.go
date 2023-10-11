package server

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ураааааа"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

}
