package Models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid      uuid.UUID `json:"uuid" db:"uuid"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Name      string    `json:"name" db:"name"`
	Surname   string    `json:"surname" db:"surname"`
	Nickname  string    `json:"nickname" db:"nickname"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	LastLogin time.Time `json:"lastLogin" db:"last_login"`
}
