package utils

import "golang.org/x/crypto/bcrypt"

func PasswordMatchesHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
