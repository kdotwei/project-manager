package handlers

import (
	"main/middleware"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register new user
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User
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

// Login
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var credentials models.User
		var user models.User

		if err := context.ShouldBindJSON(&credentials); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Look for the user
		if err := db.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
			return
		}

		// Verify password
		if !user.CheckPassword(credentials.Password) {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
			return
		}

		// Generate JWT
		token, err := middleware.GenerateToken(user.Username)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Steup cookie
		context.SetCookie("token", token, 3600, "/", "", false, true)

		context.JSON(http.StatusOK, gin.H{"message": "login successful"})
	}
}

func Logout() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.SetCookie("token", "", -1, "/", "", false, true)
		context.Redirect(http.StatusTemporaryRedirect, "/auth/login")
	}
}

// ValidateToken is used by the Nginx auth_request module to validate JWT tokens
func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extract the JWT token from the Authorization header
		tokenString, err := context.Cookie("token")

		// If the token is not present, return an unauthorized status
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify the token
		_, valid := middleware.VerifyToken(tokenString)
		if !valid {
			// If the token is invalid, return an unauthorized status
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the token is valid, return OK status
		context.Status(http.StatusOK)
	}
}
