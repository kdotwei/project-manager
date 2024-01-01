package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task models.Task
		if err := context.ShouldBindJSON(&task); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if task.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Task name cannot be empty"})
			return
		}

		// Check if task name already exists
		var existingTask models.Task
		if err := db.Where("name = ? AND project_id = ?", task.Name, task.ProjectID).First(&existingTask).Error; err == nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Task name already exists in the project"})
			return
		}

		if err := db.Create(&task).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, task)
	}
}
