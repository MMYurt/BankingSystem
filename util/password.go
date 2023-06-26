package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - returns the hash of the pw
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash pw: %w", err)
	}

	return string(hashedPassword), nil
}

// CheckPassword - checks if the provided pw is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
