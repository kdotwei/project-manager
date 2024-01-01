// handlers/create_user.go
package handlers

import (
	"log"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserJSON(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User

		log.Println("Creating user")

		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Check if the user already exists
		if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "The user already exists"})
			return
		}

		// Encrypt password
		if err := user.SetPassword(user.Password); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
			return
		}

		// Saving user info
		if err := db.Create(&user).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
			return
		}

		// Get or create the 'user' role
		var role models.Role
		err := db.Where(models.Role{Name: "user"}).FirstOrCreate(&role).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign role"})
			return
		}

		// Add role to user
		if err := models.AddRoleToUser(db, &user, &role); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add role to user"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	}
}
