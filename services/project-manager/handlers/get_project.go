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
		result := db.Find(&projects)
		if result.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		for i, project := range projects {
			var tasks []models.Task
			db.Where("project_id = ?", project.ID).Find(&tasks)
			projects[i].Tasks = tasks
		}

		context.JSON(http.StatusOK, projects)
	}
}
