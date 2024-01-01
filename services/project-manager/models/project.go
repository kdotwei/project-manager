package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Project struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	ProjectID int    `json:"project_id"`
}

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
