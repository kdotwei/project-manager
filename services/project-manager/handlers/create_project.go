package handlers

import (
	"main/models"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init(gdb *gorm.DB) {
	db = gdb
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Project name cannot be empty", http.StatusBadRequest)
		return
	}

	project := models.Project{Name: name}
	result := db.Create(&project)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
}
