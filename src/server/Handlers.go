package server

import (
	"cells-auth-server/src/dto"
	"cells-auth-server/src/repository"
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	e := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	json.NewEncoder(w).Encode(e)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		writeError(w, err, 401)
		return
	}

	session, err := repository.Login(data.Email, data.Password)
	if err != nil {
		writeError(w, err, 401)
		return
	}

	err = writeJson(w, session)
	if err != nil {
		writeError(w, err, 401)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

}
