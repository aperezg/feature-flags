package project

import "time"

// Project the representation of model Project
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    int       `json:"status"`
}

// Repository the interface used to enable callings to dbs
type Repository interface {
	CreateProject(name string) error
	FindByName(name string, project *Project) error
}
