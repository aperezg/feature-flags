package project

import "github.com/pkg/errors"

type Service interface {
	CreateProject(name string) (Project, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateProject(name string) (Project, error) {
	err := s.repository.CreateProject(name)
	if err != nil {
		return Project{}, errors.Wrap(err, "The project could not be created")
	}
	return Project{}, nil
}
