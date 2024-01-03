// handlers/create_task.go
package handlers

import (
	"log"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task models.Task

		projectID, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}
		task.ProjectID = uint(projectID)

		if err := context.ShouldBindJSON(&task); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		log.Println("task:", task.ProjectID)

		var project models.Project
		if err := db.First(&project, task.ProjectID).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		task.Status = "todo"

		if err := db.Create(&task).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": task})
	}
}
