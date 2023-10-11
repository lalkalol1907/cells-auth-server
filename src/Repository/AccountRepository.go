package Repository

import (
	"cells-auth-server/src/DB"
	"cells-auth-server/src/Models"
	"cells-auth-server/src/Redis"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

func generateToken() uuid.UUID {
	return uuid.New()
}

func Login(email string, password string) (*Models.AuthSession, error) {
	accessToken := generateToken()
	refreshToken := generateToken()

	userUuid := uuid.New() // Костыль временный

	session := &Models.AuthSession{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserUuid:     userUuid,
	}

	jsonSession, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.Set(context.Background(), "session:"+accessToken.String(), accessToken.String(), time.Hour*48).Err()
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.HSet(context.Background(), "sessions", accessToken.String(), jsonSession).Err()
	if err != nil {
		return nil, err
	}

	err = DB.DB.Update("last_login", Models.User{LastLogin: time.Now()}).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetUserBySession(accessToken uuid.UUID) (*Models.User, error) {
	sessionExists, err := Redis.RedisClient.Get(context.Background(), "session:"+accessToken.String()).Result()
	if err != nil {
		return nil, err
	}
	if sessionExists == "" {
		return nil, nil
	}

	useString, err := Redis.RedisClient.HGet(context.Background(), "sessions", accessToken.String()).Result()
	if err != nil {
		return nil, err
	}

	var session *Models.AuthSession

	err = json.Unmarshal([]byte(useString), &session)
	if err != nil {
		return nil, err
	}

	var user *Models.User

	err = DB.DB.Model(&Models.User{Uuid: session.UserUuid}).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllSessions() ([]*Models.AuthSession, error) {
	sessions, err := Redis.RedisClient.HGetAll(context.Background(), "sessions").Result()
	if err != nil {
		return nil, err
	}

	var authSessions []*Models.AuthSession

	for _, sessionString := range sessions {
		var session *Models.AuthSession
		err = json.Unmarshal([]byte(sessionString), &session)
		if err != nil {
			return nil, err
		}
		authSessions = append(authSessions, session)
	}

	return authSessions, nil
}

func UpdateSession(refreshToken uuid.UUID) (*Models.AuthSession, error) {
	sessions, err := GetAllSessions()
	if err != nil {
		return nil, err
	}

	var session *Models.AuthSession

	for _, e := range sessions {
		if e.RefreshToken == refreshToken {
			session = e
			break
		}
	}

	if session == nil {
		return nil, nil
	}

	accessToken := generateToken()
	refreshToken = generateToken()

	_, err = Redis.RedisClient.Del(context.Background(), "session"+session.AccessToken.String()).Result()
	if err != nil {
		return nil, err
	}

	_, err = Redis.RedisClient.HDel(context.Background(), "sessions", session.AccessToken.String()).Result()
	if err != nil {
		return nil, err
	}

	session.RefreshToken = refreshToken
	session.AccessToken = accessToken

	jsonSession, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.Set(context.Background(), "session:"+accessToken.String(), accessToken.String(), time.Hour*48).Err()
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.HSet(context.Background(), "sessions", accessToken.String(), jsonSession).Err()
	if err != nil {
		return nil, err
	}

	return session, nil
}
