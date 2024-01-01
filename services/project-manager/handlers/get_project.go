package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"main/models"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func getProjects(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM projects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		project.Tasks = getTasks(db, project.ID)
		projects = append(projects, project)
	}

	json.NewEncoder(w).Encode(projects)
}

func getTasks(db *sql.DB, projectID int) []models.Task {
	rows, err := db.Query("SELECT id, name, status, project_id FROM tasks WHERE project_id = ?", projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.ProjectID)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}
