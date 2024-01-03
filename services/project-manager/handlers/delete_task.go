// handlers/delete_task.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		taskID, err := strconv.Atoi(context.Param("taskId"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var task models.Task
		if err := db.First(&task, taskID).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if err := db.Delete(&task).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	}
}
