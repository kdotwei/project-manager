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

		if err := db.Create(&task).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, task)
	}
}
