// handlers/get_project.go
package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProjects(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var projects []models.Project

		if err := db.Preload("Tasks").Find(&projects).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving projects"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"projects": projects})
	}
}
