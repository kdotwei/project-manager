package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var project models.Project
		if err := context.ShouldBindJSON(&project); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if project.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Project name cannot be empty"})
			return
		}

		// Check if project name already exists
		if err := db.Where("name = ?", project.Name).First(&project).Error; err == nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Project name already exists"})
			return
		}

		// Create project
		if err := db.Create(&project).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating project"})
			return
		}
		context.JSON(http.StatusCreated, project)
	}
}
