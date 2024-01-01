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
	dsn := "host=db user=admin dbname=app password=asdhjkhg85ygfvd14e7bjh port=5432 sslmode=disable"
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
	db := setupDatabase()
	CreateAdminUser(db) // Seeding

	// Routes
	adminRoutes := service.Group("/admin").Use(middleware.RequireAdminRole(db))
	{
		// Protected routes for admin
		adminRoutes.GET("/users", handlers.ListUsers(db))
		adminRoutes.GET("/users/:id", handlers.GetUser(db))
	}

	service.Run()
}