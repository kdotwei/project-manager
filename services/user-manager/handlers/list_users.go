// handlers/list_user.go
package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
			return
		}

		// Remove password field from response
		for i := range users {
			users[i].Password = ""
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}
