// handlers/delete_project.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProject(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the ID frpm URL
		projectID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		// Check if the project exists
		var project models.Project
		if err := db.First(&project, projectID).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		// Delete the project
		if err := db.Select("Tasks").Delete(&project).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting project and tasks"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Project and tasks deleted successfully"})
	}
}
