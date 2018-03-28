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
	Status    Status    `json:"status"`
}

// ProjectStatus define the states in which our project can be found
type Status int

const (
	StatusDisabled Status = iota
	StatusEnabled
)

// Repository provides access a Project store.
type Repository interface {
	FindByName(name string) (Project, error)
	FindByID(ID Identity) (Project, error)
	Persist(project *Project) error
}
