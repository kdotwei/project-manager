// handlers/create_project.go
package handlers

import (
	"log"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var project models.Project

		log.Println("Creating project")
		if err := context.ShouldBindJSON(&project); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		log.Println("Project Name: ", project.Name)

		// Check if the project exists
		if err := db.Where("name = ?", project.Name).First(&models.Project{}).Error; err == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "The project already exists"})
			return
		}

		// Save the project
		if err := db.Create(&project).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Create project failed"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Create project successful"})
	}
}
