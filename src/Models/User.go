package Models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid      uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Surname   string    `json:"surname"`
	Nickname  string    `json:"nickname" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	LastLogin time.Time `json:"lastLogin"`
}
