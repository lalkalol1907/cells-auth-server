package Handlers

import (
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/Repository"
	"cells-auth-server/src/Server/HttpTools"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	session, err := Repository.Login(data.Email, data.Password)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	err = HttpTools.WriteJson(w, session)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
	}
}
