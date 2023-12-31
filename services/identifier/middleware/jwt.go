package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

// IsLoggedIn checks if the user is logged in by validating the JWT token
func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token") // Assuming token is stored in a cookie
		if err != nil {
			// Token not found in cookie, redirect to login
			c.Redirect(http.StatusTemporaryRedirect, "/identifier/login")
			c.Abort()
			return
		}

		// Verify the token
		claims, valid := VerifyToken(tokenString)
		if !valid {
			// Token is invalid, redirect to login
			c.Redirect(http.StatusTemporaryRedirect, "/identifier/login")
			c.Abort()
			return
		}

		// Token is valid, set the username in context and continue
		c.Set("username", claims.Username)
		c.Next()
	}
}
