package main

import (
	"fmt"
	"log"
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

	db.AutoMigrate(&models.User{}, &models.Role{})

	return db
}

func CreateAdminUser(db *gorm.DB) {
	// Check if the role admin exists
	var adminRole models.Role
	if err := db.Where("name = ?", "admin").FirstOrCreate(&adminRole, models.Role{Name: "admin"}).Error; err != nil {
		log.Fatalf("Failed to create admin role: %v", err)
	}

	// Check if the user admin exists
	var adminUser models.User
	err := db.Where("username = ?", "admin").First(&adminUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatalf("Failed to check if admin user exists: %v", err)
	}

	// If not, create admin user
	if err == gorm.ErrRecordNotFound {
		adminUser = models.User{
			Username: "admin",
			Password: "admin", // TODO: set a stronger password
			Roles:    []models.Role{adminRole},
		}
		adminUser.SetPassword(adminUser.Password)

		if err := db.Create(&adminUser).Error; err != nil {
			log.Fatalf("Failed to create admin user: %v", err)
		}

		fmt.Println("Admin user created")
	}
}

func main() {
	// Initialize
	service := gin.Default()
	service.LoadHTMLGlob("templates/html/*")
	service.Static("/assets", "./templates/assets")
	db := setupDatabase()

	// Seeding
	CreateAdminUser(db)

	// Routes setting
	apiRoutes := service.Group("/api").Use(middleware.RequireAdminRole(db))
	{
		// API requests
		apiRoutes.GET("/users", handlers.ListUsers(db))
		apiRoutes.GET("/users/:id", handlers.GetUser(db))

		apiRoutes.POST("/users/create", handlers.CreateUserJSON(db))
		apiRoutes.PUT("/users/:id/update", handlers.UpdateUser(db))
		apiRoutes.DELETE("/users/:id/delete", handlers.DeleteUser(db))
	}

	adminRoutes := service.Group("/").Use(middleware.RequireAdminRole(db))
	{
		// Page requests
		adminRoutes.GET("/users", handlers.IndexPage)
		adminRoutes.GET("/users/:id", handlers.UserPage)
		adminRoutes.GET("/users/create", handlers.CreatePage)
		adminRoutes.GET("/users/:id/edit", handlers.EditPage)
	}

	service.Run()
}
