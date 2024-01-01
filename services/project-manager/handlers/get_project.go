package handlers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	result := db.Preload("Tasks").Find(&projects)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}
