package DTO

import "github.com/google/uuid"

type RefreshTokenDto struct {
	RefreshToken uuid.UUID `json:"refresh_token"`
}
