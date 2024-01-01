package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("ashf23irj84yt0-jka")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Generate JWT
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// Verify JWT
func VerifyToken(tokenString string) (*Claims, bool) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, false
	}

	return claims, token.Valid
}
