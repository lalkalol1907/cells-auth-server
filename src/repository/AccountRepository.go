package repository

import (
	"cells-auth-server/src/models"
	"cells-auth-server/src/redis"
	"context"
	"encoding/json"
	"github.com/google/uuid"
)

func generateToken() uuid.UUID {
	return uuid.New()
}

func Login(email uuid.UUID, password string) (*models.AuthSession, error) {
	accessToken := generateToken()
	refreshToken := generateToken()

	userUuid := email // Костыль временный

	session := &models.AuthSession{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserUuid:     userUuid,
	}

	jsonSession, err := json.Marshal(session)

	if err != nil {
		return nil, err
	}

	err = redis.RedisClient.HSet(context.Background(), "sessions", accessToken, string(jsonSession)).Err()

	if err != nil {
		return nil, err
	}

	return session, nil
}
