// handlers/get_tasks.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProjectTasks(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the project ID
		projectID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		// Get tasks for the project
		var project models.Project
		if err := db.Preload("Tasks").First(&project, projectID).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"tasks": project.Tasks})
	}
}
