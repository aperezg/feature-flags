package project

import "gopkg.in/src-d/go-errors.v1"

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
	var errorCreatingProject = errors.NewKind("The project could not be created")

	err := s.repository.CreateProject(name)
	if err != nil {
		return Project{}, errorCreatingProject.Wrap(err)
	}
	return Project{}, nil
}
