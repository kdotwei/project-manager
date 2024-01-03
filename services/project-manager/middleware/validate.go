package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequireLogin(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString, err := context.Cookie("token")
		if err != nil {
			// token not found
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Redirect(http.StatusTemporaryRedirect, "/auth/logout")
			return
		}

		// Validate token
		_, valid := VerifyToken(tokenString)
		if !valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Redirect(http.StatusTemporaryRedirect, "/auth/logout")
			return
		}

		context.Next()
	}
}
