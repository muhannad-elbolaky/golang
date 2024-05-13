package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hashed password from the given plaintext password.
//
// It takes a plaintext password as input and returns the hashed password and an error.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if the provided password matches the hashed password.
//
// Parameters:
// - password: the plaintext password to check.
// - hashedPassword: the hashed password to compare against.
//
// Returns:
// - bool: true if the password matches the hashed password, false otherwise.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
