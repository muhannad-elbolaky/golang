package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "basmala"

// GenerateToken generates a JWT token with the given email and userID.
//
// Parameters:
// - email: a pointer to a string representing the user's email.
// - userID: a pointer to an int64 representing the user's ID.
//
// Returns:
// - string: the generated JWT token.
// - error: an error if the token generation fails.
func GenerateToken(email *string, userID *int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates the given JWT token and extracts the user ID from it.
//
// Parameters:
// - token: a string representing the JWT token to be validated.
//
// Returns:
// - int64: the user ID extracted from the token.
// - error: an error if the token validation fails.
func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsVaild := parsedToken.Valid
	if !tokenIsVaild {
		return 0, errors.New("invalid token")

	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("unexpected signing method")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil
}
