package Repository

import (
	"cells-auth-server/src/CustomErrors"
	"cells-auth-server/src/DB"
	"cells-auth-server/src/Models"
	"cells-auth-server/src/Redis"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"time"
)

func generateToken() string {
	return uuid.New().String()
}

func CreateSession(userUuid uuid.UUID) (*Models.AuthSession, error) {
	accessToken := generateToken()
	refreshToken := generateToken()

	session := &Models.AuthSession{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserUuid:     userUuid,
	}

	jsonSession, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.Set(context.Background(), "session:"+accessToken, accessToken, time.Hour*48).Err()
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.HSet(context.Background(), "sessions", accessToken, jsonSession).Err()
	if err != nil {
		return nil, err
	}

	DB.DB.QueryRow(context.Background(), "UPDATE users SET last_login=$1 WHERE uuid=$2", time.Now(), userUuid)

	return session, nil
}

func GetUserByEmail(email string) (uuid.UUID, string, error) {
	var userUuid uuid.UUID
	var password string

	err := DB.DB.QueryRow(
		context.Background(),
		"SELECT uuid, password FROM users WHERE email=$1 LIMIT 1",
		email,
	).Scan(&userUuid, &password)
	return userUuid, password, err
}

func CreateUser(email string, password string, name string, surname string, nickname string) (uuid.UUID, error) {
	var userUUID uuid.UUID

	err := DB.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (email, password, name, surname, nickname) VALUES ($1, $2, $3, $4, $5) RETURNING uuid;",
		email,
		password,
		name,
		surname,
		nickname,
	).Scan(&userUUID)

	return userUUID, err
}

func GetUserBySession(accessToken uuid.UUID) (*Models.User, error) {
	useString, err := Redis.RedisClient.HGet(context.Background(), "sessions", accessToken.String()).Result()
	if err == redis.Nil {
		return nil, CustomErrors.NoSession
	}
	if err != nil {
		return nil, err
	}

	_, err = Redis.RedisClient.Get(context.Background(), "session:"+accessToken.String()).Result()
	if err == redis.Nil {
		return nil, CustomErrors.NeedRefreshError
	}

	var session *Models.AuthSession

	err = json.Unmarshal([]byte(useString), &session)
	if err != nil {
		return nil, err
	}

	rows, err := DB.DB.Query(
		context.Background(),
		"SELECT * FROM users WHERE uuid=$1 LIMIT 1",
		session.UserUuid,
	)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[Models.User])
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

func DeleteSession(accessToken string) error {
	_, err := Redis.RedisClient.Del(context.Background(), "session"+accessToken).Result()
	if err != nil {
		return err
	}

	_, err = Redis.RedisClient.HDel(context.Background(), "sessions", accessToken).Result()
	if err != nil {
		return err
	}

	return nil
}
