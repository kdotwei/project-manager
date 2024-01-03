package models

import (
	"gorm.io/gorm"
)

// Project has many tasks
type Project struct {
	gorm.Model
	Name  string `json:"name"`
	Tasks []Task `gorm:"foreignKey:ProjectID" json:"tasks"` // Foreign key
}

// Task belongs to a project
type Task struct {
	gorm.Model
	Name      string  `json:"name"`
	Status    string  `json:"status"`
	ProjectID uint    `json:"project_id"` // Foreign key
	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

// CreateProjectWithTasks is responsible for creating a new project with tasks
func CreateProjectWithTasks(db *gorm.DB, project *Project) error {
	return db.Create(project).Error
}

// FindProjectWithTasks is responsible for finding a project by tasks
func FindProjectWithTasks(db *gorm.DB, projectID uint) (*Project, error) {
	var project Project
	err := db.Preload("Tasks").First(&project, projectID).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}
