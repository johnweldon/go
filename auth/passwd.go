package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func TestPassword(possible string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(possible))
	return err == nil
}
