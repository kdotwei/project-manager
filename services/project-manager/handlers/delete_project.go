package handlers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
			return
		}

		// Look up the project
		var project models.Project
		if err := db.First(&project, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		// Delete tasks associated with the project
		db.Where("project_id = ?", id).Delete(&models.Task{})

		// Delete the project
		db.Delete(&project)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
