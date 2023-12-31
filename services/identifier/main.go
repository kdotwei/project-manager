package main

import (
	"fmt"
	"main/handlers"
	"main/models"
	"net/http"

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

	// Migration
	db.AutoMigrate(&models.User{})

	return db
}

func LoginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	// Serivce initialization
	service := gin.Default()
	service.LoadHTMLGlob("templates/html/*")
	service.Static("/assets", "./templates/assets")
	db := setupDatabase()

	// Route
	service.GET("/login", LoginPage)

	// From Page
	service.POST("/login", handlers.Login(db))
	service.POST("/register", handlers.Register(db))

	service.Run() // listen and serve on 0.0.0.0:8080
}
