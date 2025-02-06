package models

import (
	"time"

	"gorm.io/gorm"
)

type App struct {
	gorm.Model

	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	AppKey    string    `gorm:"unique" json:"app_key"`
	AppSecret string    `gorm:"unique" json:"app_secret"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (App) TableName() string {
	return "apps"
}
