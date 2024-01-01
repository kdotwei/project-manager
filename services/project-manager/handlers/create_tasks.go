package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	_, err = db.Exec("INSERT INTO tasks (name, status, project_id) VALUES (?, ?, ?)", name, "Pending", projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
