package handlers

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func CreateProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Project name cannot be empty", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
