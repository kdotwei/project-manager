package main

import (
	"database/sql"
	"log"
	"net/http"

	"main/handlers"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "projects.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProjects(db, w, r)
	})

	http.HandleFunc("/createProject", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateProject(db, w, r)
	})

	http.HandleFunc("/createTask", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(db, w, r)
	})

	log.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}

func createTable() {
	createProjectsTableSQL := `
		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`
	createTasksTableSQL := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			status TEXT,
			project_id INTEGER,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		);
	`
	_, err := db.Exec(createProjectsTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createTasksTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
