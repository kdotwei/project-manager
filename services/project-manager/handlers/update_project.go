// handlers/update_project.go
package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProject(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		projectID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		var updateInfo struct {
			Name string `json:"name"`
		}

		if err := context.ShouldBindJSON(&updateInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		if err := db.Model(&models.Project{}).Where("id = ?", projectID).Update("name", updateInfo.Name).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating project"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	}
}
