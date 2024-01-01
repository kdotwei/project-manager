// handlers/delete_user.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		idParam := context.Param("id")
		userID, err := strconv.Atoi(idParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
