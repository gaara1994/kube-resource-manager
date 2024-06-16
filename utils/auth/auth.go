package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword  takes a plain text password and returns its bcrypt hash.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// ComparePasswords  compares a plain text password with a bcrypt hashed password.
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
