package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, expire int64) (string, error) {
	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": expire,
	})

	// Sign Token
	tokenString, signTokenErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if signTokenErr != nil {
		return "", signTokenErr
	}

	return tokenString, nil
}
