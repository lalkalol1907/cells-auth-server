package Handlers

import (
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/HttpServer/HttpTools"
	"cells-auth-server/src/Repository"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.RegisterDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	_, _, err = Repository.GetUserByEmail(data.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	if err == nil {
		HttpTools.WriteError(w, errors.New("user already exists"), http.StatusInternalServerError)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	userUuid, err := Repository.CreateUser(data.Email, string(hashed), data.Name, data.Surname, data.Nickname)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	session, err := Repository.CreateSession(userUuid)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = HttpTools.WriteJson(w, session)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
	}
}
