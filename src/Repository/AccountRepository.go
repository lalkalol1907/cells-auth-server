package Repository

import (
	"cells-auth-server/src/DB"
	"cells-auth-server/src/DTO"
	"cells-auth-server/src/Models"
	"cells-auth-server/src/Redis"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func generateToken() uuid.UUID {
	return uuid.New()
}

func createSession(userUuid uuid.UUID) (*Models.AuthSession, error) {
	fmt.Print("Вызывлся")
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

	err = Redis.RedisClient.Set(context.Background(), "session:"+accessToken.String(), accessToken.String(), time.Hour*48).Err()
	if err != nil {
		return nil, err
	}

	err = Redis.RedisClient.HSet(context.Background(), "sessions", accessToken.String(), jsonSession).Err()
	if err != nil {
		return nil, err
	}

	DB.DB.QueryRow(context.Background(), "UPDATE users SET last_login=$1 WHERE uuid=$2", time.Now(), userUuid)

	return session, nil
}

func Login(dto *DTO.LoginDto) (*Models.AuthSession, error) {

	var uuid uuid.UUID
	var password string

	err := DB.DB.QueryRow(
		context.Background(),
		"SELECT uuid, password FROM users WHERE email=$1 LIMIT 1",
		dto.Email,
	).Scan(&uuid, &password)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user doesn't exist")
	}
	if err != nil {
		return nil, err
	}

	notSuccess := bcrypt.CompareHashAndPassword([]byte(password), []byte(dto.Password))
	if notSuccess != nil {
		return nil, errors.New("incorrect password")
	}

	return createSession(uuid)
}

func Register(dto *DTO.RegisterDto) (*Models.AuthSession, error) {
	err := DB.DB.QueryRow(
		context.Background(),
		"SELECT uuid FROM users WHERE email=$1 LIMIT 1",
		dto.Email,
	).Scan(nil)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {

		return nil, err
	}

	var userUUID uuid.UUID

	err = DB.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (email, password, name, surname, nickname) VALUES ($1, $2, $3, $4, $5) RETURNING uuid;",
		dto.Email,
		string(hashed),
		dto.Name,
		dto.Surname,
		dto.Nickname,
	).Scan(&userUUID)

	if err != nil {
		return nil, err
	}

	return createSession(userUUID)
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

	_, err = Redis.RedisClient.Del(context.Background(), "session"+session.AccessToken.String()).Result()
	if err != nil {
		return nil, err
	}

	_, err = Redis.RedisClient.HDel(context.Background(), "sessions", session.AccessToken.String()).Result()
	if err != nil {
		return nil, err
	}

	return createSession(session.UserUuid)
}
