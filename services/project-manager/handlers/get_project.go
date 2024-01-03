// handlers/get_project.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProject(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the ID from URL
		projectID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		var project models.Project

		// Find the project ID from database
		if err := db.Preload("Tasks").First(&project, projectID).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		context.JSON(http.StatusOK, project)
	}
}
