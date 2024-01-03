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
	dsn := "host=db user=admin dbname=app password=admin port=5432 sslmode=disable"
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
		apiRoutes.GET("/projects", handlers.ListProjects(db))
		apiRoutes.GET("/projects/:id", handlers.GetProject(db))
		apiRoutes.POST("/projects/create", handlers.CreateProject(db))
		apiRoutes.PUT("/projects/:id/update", handlers.UpdateProject(db))
		apiRoutes.DELETE("/projects/:id/delete", handlers.DeleteProject(db))

		// CRUD for tasks
		apiRoutes.GET("/projects/:id/tasks", handlers.ListProjectTasks(db))
		apiRoutes.GET("/projects/:id/tasks/:taskId", handlers.GetTask(db))
		apiRoutes.POST("/projects/:id/tasks/create", handlers.CreateTask(db))
		apiRoutes.PUT("/projects/:id/tasks/:taskId/update", handlers.UpdateTask(db))
		apiRoutes.DELETE("/projects/:id/tasks/:taskId/delete", handlers.DeleteTask(db))
	}

	loginRoutes := service.Group("/").Use(middleware.RequireLogin(db))
	{
		// Pages for projects
		loginRoutes.GET("/projects", handlers.IndexPage)
		loginRoutes.GET("/projects/:id/edit", handlers.EditPage)

		// Pages for tasks
		loginRoutes.GET("/projects/:id/tasks", handlers.ProjectPage)
		loginRoutes.GET("/projects/:id/tasks/:taskId/edit", handlers.EditTaskPage)
	}

	service.Run()
}
