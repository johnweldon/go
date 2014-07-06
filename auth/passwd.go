package auth

import (
	"code.google.com/p/go.crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func TestPassword(possible string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(possible))
	return err == nil
}
