package db

import (
	"log"
	"os"

	"github.com/socketbasehq/socketbase/pkg/models"
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

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.App{})

	CreateAdminUser()
}

func CreateAdminUser() {
	user := models.User{
		Username: "admin",
		Password: "admin",
	}

	if DB.Where("username = ?", user.Username).First(&user).Error == nil {
		return
	}

	err := DB.Create(&user).Error
	if err != nil {
		log.Fatal("Failed to create admin user:", err)
	}
}
