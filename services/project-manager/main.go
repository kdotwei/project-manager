package main

import (
	"log"
	"net/http"

	"main/handlers"
	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("projects.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Project{}, &models.Task{})
	handlers.Init(db)

	http.HandleFunc("/projects", handlers.GetProjects)

	http.HandleFunc("/createProject", handlers.CreateProject)

	http.HandleFunc("/createTask", handlers.CreateTask)

	log.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}
