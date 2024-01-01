package middleware

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequireAdminRole(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString, err := context.Cookie("token")
		if err != nil {
			// token not found
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Validate token
		claims, valid := VerifyToken(tokenString)
		if !valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check if the user is admin
		user := models.User{}
		if err := db.Where("username = ?", claims.Username).First(&user).Error; err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		if !models.UserHasRole(db, &user, "admin") {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Requires admin role"})
			return
		}

		context.Next()
	}
}
