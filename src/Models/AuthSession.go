package Models

import (
	"github.com/google/uuid"
)

type AuthSession struct {
	AccessToken  uuid.UUID `json:"accessToken"`
	RefreshToken uuid.UUID `json:"refreshToken"`
	UserUuid     uuid.UUID `json:"userUuid"`
}
