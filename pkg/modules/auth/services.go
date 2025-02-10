package auth

import (
	"github.com/socketbasehq/socketbase/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CreateAdminUser() {
	db := db.DB

	user := models.User{
		Username: "admin",
		Password: hashPassword("admin"),
	}

	if db.Where("username = ?", user.Username).First(&user).Error == nil {
		return
	}

	db.Create(&user)
}
