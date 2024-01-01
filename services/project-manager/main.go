package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "projects.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
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

func main() {
	defer db.Close()

	http.HandleFunc("/projects", handlers.getProjects(db))
	http.HandleFunc("/createProject", handlers.createProject(db))
	http.HandleFunc("/createTask", handlers.createTask(db))
	log.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}
