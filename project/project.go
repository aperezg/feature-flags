package project

import (
	. "github.com/aperezg/feature-flags/identity"
	"time"
)

// Project the representation of model Project
type Project struct {
	ID        Identity  `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    int       `json:"status"`
}

// Repository the interface used to enable callings to dbs
type Repository interface {
	FindByName(name string) (Project, error)
	FindByID(ID Identity) (Project, error)
	Persist(project *Project) error
}
