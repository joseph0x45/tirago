package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func PasswordMatchesHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to generate password hash: %w", err)
	}
	return string(hash), nil
}
