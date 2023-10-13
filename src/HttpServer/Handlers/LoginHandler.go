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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data DTO.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusBadRequest)
		return
	}

	userUuid, password, err := Repository.GetUserByEmail(data.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		HttpTools.WriteError(w, errors.New("user not found"), http.StatusUnauthorized)
		return
	}
	if err != nil {
		HttpTools.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password))
	if err != nil {
		HttpTools.WriteError(w, errors.New("incorrect password"), http.StatusUnauthorized)
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
