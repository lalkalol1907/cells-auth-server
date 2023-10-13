package Handlers

import (
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/HttpServer/HttpTools"
	"cells-auth-server/src/Models"
	"cells-auth-server/src/Repository"
	"encoding/json"
	"errors"
	"net/http"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.RefreshTokenDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	sessions, err := Repository.GetAllSessions()
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	var session *Models.AuthSession

	for _, e := range sessions {
		if e.RefreshToken == data.RefreshToken {
			session = e
			break
		}
	}
	if session == nil {
		HttpTools.WriteError(w, errors.New("no refresh token"), http.StatusUnauthorized)
		return
	}

	err = Repository.DeleteSession(session.AccessToken)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	newSession, err := Repository.CreateSession(session.UserUuid)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = HttpTools.WriteJson(w, newSession)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
	}
}
