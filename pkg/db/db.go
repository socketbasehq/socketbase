package db

import (
	"log"
	"os"

	"github.com/socketbase/socketbase/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
