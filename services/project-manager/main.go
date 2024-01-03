package main

import (
	"fmt"

	"main/handlers"
	"main/middleware"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupDatabase() *gorm.DB {
	dsn := "host=db user=admin dbname=app password=asdhjkhg85ygfvd14e7bjh port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&models.Project{}, &models.Task{})

	return db
}

func main() {
	service := gin.Default()
	service.LoadHTMLGlob("templates/html/*")
	service.Static("/assets", "./templates/assets")
	db := setupDatabase()

	apiRoutes := service.Group("/api").Use(middleware.RequireLogin(db))
	{
		apiRoutes.GET("/projects", handlers.GetProjects(db))
		apiRoutes.POST("/createProject", handlers.CreateProject(db))
		apiRoutes.POST("/createTask", handlers.CreateTask(db))
	}
	{
		service.GET("/projects", handlers.IndexPage)
	}
}
