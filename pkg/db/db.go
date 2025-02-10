package db

import (
	"log"

	"github.com/socketbasehq/socketbase/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.App{})
}
