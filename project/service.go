package project

import "github.com/pkg/errors"

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
	err := s.repository.CreateProject(name)
	if err != nil {
		return Project{}, errors.Wrap(err, "The project could not be created")
	}
	var p Project
	err = s.repository.FindByName(name, &p)
	if err != nil {
		return Project{}, errors.Wrap(err, "The current created project could not be found")
	}
	return p, nil
}
