package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var projects []models.Project
		result := db.Preload("Tasks").Find(&projects)
		if result.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		context.JSON(http.StatusOK, projects)
	}
}
