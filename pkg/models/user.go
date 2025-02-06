package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Username  string    `gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
