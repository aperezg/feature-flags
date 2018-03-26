package project

import "time"

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    int       `json:"status"`
}

type Repository interface {
	CreateProject(name string) error
}
