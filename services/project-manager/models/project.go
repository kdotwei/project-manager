package models

type Project struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks" gorm:"foreignKey:ProjectID"`
}

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	ProjectID int    `json:"project_id"`
}
