// handlers/create_user.go
package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var newUser models.User
		if err := context.ShouldBindJSON(&newUser); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// Password
		if err := newUser.SetPassword(newUser.Password); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
			return
		}

		// Create user in database
		if err := db.Create(&newUser).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		newUser.Password = ""
		context.JSON(http.StatusCreated, gin.H{"user": newUser})
	}
}
