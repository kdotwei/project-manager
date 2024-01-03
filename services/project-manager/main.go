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
		// CRUD for projects
		apiRoutes.GET("/projects", handlers.GetProjects(db))
		apiRoutes.POST("/projects/create", handlers.CreateProject(db))
		apiRoutes.PUT("/projects/:id/update")
		apiRoutes.DELETE("/projects/:id/delete")

		// CRUD for tasks
		apiRoutes.GET("/projects/:id/tasks")
		apiRoutes.POST("/projects/:id/tasks/create")
		apiRoutes.PUT("/projects/:id/tasks/:id/update")
		apiRoutes.DELETE("/projects/:id/tasks/:id/delete")
	}

	loginRoutes := service.Group("/").Use(middleware.RequireLogin(db))
	{
		// Pages for projects
		loginRoutes.GET("/projects", handlers.IndexPage)
		loginRoutes.GET("/projects/:id/edit")

		// Pages for tasks
		loginRoutes.GET("/projects/:id/tasks")
		loginRoutes.GET("/projects/:id/tasks/:id/edit")
	}

	service.Run()
}
