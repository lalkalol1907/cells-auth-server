package Repository

import (
	"cells-auth-server/src/Models"
	"cells-auth-server/src/Redis"
	"context"
	"encoding/json"
	"github.com/google/uuid"
)

func generateToken() uuid.UUID {
	return uuid.New()
}

func Login(email uuid.UUID, password string) (*Models.AuthSession, error) {
	accessToken := generateToken()
	refreshToken := generateToken()

	userUuid := email // Костыль временный

	session := &Models.AuthSession{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserUuid:     userUuid,
	}

	jsonSession, err := json.Marshal(session)

	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.HSet(context.Background(), "sessions", accessToken, string(jsonSession)).Err()

	if err != nil {
		return nil, err
	}

	return session, nil
}
