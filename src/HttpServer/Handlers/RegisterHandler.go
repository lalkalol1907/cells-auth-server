package Handlers

import (
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/HttpServer/HttpTools"
	"cells-auth-server/src/Repository"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.RegisterDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	session, err := Repository.Register(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	err = HttpTools.WriteJson(w, session)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
	}
}
