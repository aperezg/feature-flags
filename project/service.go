package project

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"time"
)

// Service the interface used for encapsulate the business logic of the project
type Service interface {
	CreateProject(name string) (Project, error)
}

type service struct {
	repository Repository
}

// NewService Initialize the service Project
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// CreateProject this function creates a new project based on the name passed by parameter
func (s *service) CreateProject(name string) (Project, error) {
	if p, _ := s.repository.FindByName(name); p != (Project{}) {
		return Project{}, errors.New("The project %s already exists")
	}

	p := Project{
		ID:        uuid.NewV4().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Status:    1,
	}

	err := s.repository.Persist(&p)
	if err != nil {
		return Project{}, errors.Wrap(err, "The project could not be created")
	}
	return p, nil
}
