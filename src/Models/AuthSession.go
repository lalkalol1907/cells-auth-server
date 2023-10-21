package Models

import (
	"github.com/google/uuid"
)

type AuthSession struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	UserUuid     uuid.UUID `json:"userUuid"`
}
