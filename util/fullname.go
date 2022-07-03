package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashFuLLName( full_name string) (string, error) {
	hashedFullName, err := bcrypt.GenerateFromPassword([]byte(full_name), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash username: %w", err)
	}
	return string(hashedFullName), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckFullName(full_name string, hashedFullName string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedFullName), []byte(full_name))
}


