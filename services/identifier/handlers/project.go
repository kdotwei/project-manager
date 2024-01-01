package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func main() {
	defer db.Close()

	http.HandleFunc("/projects", getProjects)
	http.HandleFunc("/createProject", createProject)
	http.HandleFunc("/createTask", createTask)
	log.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM projects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		project.Tasks = getTasks(project.ID)
		projects = append(projects, project)
	}

	json.NewEncoder(w).Encode(projects)
}

func getTasks(projectID int) []Task {
	rows, err := db.Query("SELECT id, name, status, project_id FROM tasks WHERE project_id = ?", projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.ProjectID)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func createProject(w http.ResponseWriter, r *http.Request) {
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

func createTask(w http.ResponseWriter, r *http.Request) {
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
