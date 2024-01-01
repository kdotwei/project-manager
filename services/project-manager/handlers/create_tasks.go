package handlers

import (
	"main/models"
	"net/http"
	"strconv"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	projectID, err := strconv.Atoi(r.FormValue("projectId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if name == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	task := models.Task{Name: name, Status: "Pending", ProjectID: projectID}
	result := db.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
