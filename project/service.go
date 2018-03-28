package project

import (
	"fmt"
	. "github.com/aperezg/feature-flags/identity"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"time"
)

// Service the interface used for encapsulate the business logic of the project
type Service interface {
	CreateProject(name string) (Project, error)
	ModifyProjectName(ID Identity, newName string) (Project, error)
	DeactivateProject(ID Identity) error
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
		ID:        Identity(uuid.NewV4().String()),
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

// ModifyProjectName First of all searching into a repository if there any project with the name to change
// if is not, search by Identity the project and when found it then change the name and Persists the changes
func (s *service) ModifyProjectName(ID Identity, newName string) (Project, error) {
	if p, _ := s.repository.FindByName(newName); p != (Project{}) {
		return Project{}, errors.New(fmt.Sprintf("The project %s already exists", newName))
	}

	p, err := s.repository.FindByID(ID)
	if err != nil {
		return Project{}, errors.Wrap(err, fmt.Sprintf("The project %s not found", ID.String()))
	}
	p.Name = newName

	err = s.repository.Persist(&p)
	if err != nil {
		return Project{}, errors.Wrap(err, "The changes to the project could not be saved")
	}

	return p, nil
}

func (s *service) DeactivateProject(ID Identity) error {
	_, err := s.repository.FindByID(ID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("The project %s not found", ID.String()))
	}
	return nil
}
