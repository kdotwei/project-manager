// handlers/update_task.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateTask(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		taskID, err := strconv.Atoi(context.Param("taskId"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var updateInfo struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		}
		if err := context.ShouldBindJSON(&updateInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		if err := db.Model(&models.Task{}).Where("id = ?", taskID).Updates(models.Task{Name: updateInfo.Name, Status: updateInfo.Status}).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	}
}
