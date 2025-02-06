package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
