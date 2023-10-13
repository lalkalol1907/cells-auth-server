package Handlers

import (
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/Repository"
	"cells-auth-server/src/Server/HttpTools"
	"encoding/json"
	"net/http"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.RefreshTokenDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	session, err := Repository.UpdateSession(data.RefreshToken)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	err = HttpTools.WriteJson(w, session)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
	}
}
