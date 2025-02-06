package apps

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

func generateRandomInt() int {
	min := 100000
	max := 999999
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return int(bigInt.Int64()) + min
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// generateAppCredentials generates a new set of app credentials
func generateAppCredentials() (id int, key, secret string, err error) {

	id = generateRandomInt()

	key, err = generateRandomString(18)
	if err != nil {
		return 0, "", "", err
	}

	secret, err = generateRandomString(20)
	if err != nil {
		return 0, "", "", err
	}

	return id, key, secret, nil
}
