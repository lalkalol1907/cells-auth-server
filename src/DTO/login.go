package DTO

import "github.com/google/uuid"

type LoginDto struct {
	Email    uuid.UUID `json:"email"`
	Password string    `json:"password"`
}
