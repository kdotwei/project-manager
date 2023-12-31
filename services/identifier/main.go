package main

import (
	"fmt"
	"main/handlers"
	"main/middleware"
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
	db.AutoMigrate(&models.User{}, &models.Role{})

	// Checking default roles
	ensureRoles(db, []string{"user", "admin"})

	return db
}

func ensureRoles(db *gorm.DB, roles []string) {
	for _, roleName := range roles {
		var role models.Role
		if err := db.Where(models.Role{Name: roleName}).FirstOrCreate(&role, models.Role{Name: roleName}).Error; err != nil {
			fmt.Printf("Unable to create role '%s': %v\n", roleName, err)
		}
	}
}

func loginPage(context *gin.Context) {
	tokenString, err := context.Cookie("token")
	if err == nil {
		_, valid := middleware.VerifyToken(tokenString)
		if valid {
			context.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
	}
	context.HTML(http.StatusOK, "login.html", nil)
}

func registerPage(context *gin.Context) {
	tokenString, err := context.Cookie("token")
	if err == nil {
		_, valid := middleware.VerifyToken(tokenString)
		if valid {
			context.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
	}
	context.HTML(http.StatusOK, "register.html", nil)
}

func main() {
	// Serivce initialization
	service := gin.Default()
	service.LoadHTMLGlob("templates/html/*")
	service.Static("/assets", "./templates/assets")
	db := setupDatabase()

	// Route
	service.GET("/login", loginPage)
	service.GET("/register", registerPage)
	service.GET("/logout", handlers.Logout())

	// From Page
	service.POST("/login", handlers.Login(db))
	service.POST("/register", handlers.Register(db))

	service.Run() // listen and serve on 0.0.0.0:8080
}
