// handlers/get_user.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		idParam := context.Param("id")
		userID, err := strconv.Atoi(idParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var user models.User
		if err := db.Preload("Roles").First(&user, userID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		user.Password = ""

		context.JSON(http.StatusOK, gin.H{"user": user})
	}
}
