package project

import (
	"fmt"
	"time"

	"github.com/aperezg/feature-flags/identity"
	"github.com/pkg/errors"
)

// Service the interface used for encapsulate the business logic of the project
type Service interface {
	// CreateProject this function creates a new project based on the name passed by parameter
	CreateProject(name string) (Project, error)

	// ModifyProjectName First of all searching into a repository if there any project with the name to change
	// if is not, search by ID the project and when found it then change the name and Persists the changes
	ModifyProjectName(ID string, newName string) (Project, error)

	// DeactivateProject First look for a project by its ID, if it finds it then it deactivates it
	DeactivateProject(ID string) error

	// ActivateProject First look for a project by its ID, if it finds it then it activates it
	ActivateProject(ID string) error

	// RemoveProject Remove a project from the face of the earth
	RemoveProject(ID string) error

	// FindProjects Fetch all projects on store
	FindProjects() ([]Project, error)
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

func (s *service) CreateProject(name string) (Project, error) {
	if p, _ := s.repository.FindByName(name); p != (Project{}) {
		return Project{}, errors.New("The project %s already exists")
	}
	p := Project{
		ID:        identity.NewID(),
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

func (s *service) ModifyProjectName(ID string, newName string) (Project, error) {
	if p, _ := s.repository.FindByName(newName); p != (Project{}) {
		return Project{}, errors.New(fmt.Sprintf("The project %s already exists", newName))
	}

	p, err := s.checkIfProjectFound(ID)
	if err != nil {
		return p, err
	}
	p.Name = newName
	p.UpdatedAt = time.Now()

	err = s.repository.Persist(&p)
	if err != nil {
		return Project{}, errors.Wrap(err, "The name change to the project could not be saved")
	}

	return p, nil
}

func (s *service) DeactivateProject(ID string) error {
	p, err := s.checkIfProjectFound(ID)
	if err != nil {
		return err
	}
	p.Status = StatusDisabled
	p.UpdatedAt = time.Now()
	err = s.repository.Persist(&p)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Can't deactivate the project %s", p.Name))
	}

	return nil
}

func (s *service) ActivateProject(ID string) error {
	p, err := s.checkIfProjectFound(ID)
	if err != nil {
		return err
	}
	p.Status = StatusEnabled
	p.UpdatedAt = time.Now()
	err = s.repository.Persist(&p)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Can't activate the project %s", p.Name))
	}
	return nil
}

func (s *service) RemoveProject(ID string) error {
	p, err := s.checkIfProjectFound(ID)
	if err != nil {
		return err
	}

	if err := s.repository.Remove(ID); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Can't remove the project %s", p.Name))
	}
	return nil
}

func (s *service) FindProjects() ([]Project, error) {
	projects, err := s.repository.FindAll()
	if err != nil {
		return []Project{}, errors.Wrap(err, "Can't retrieve the results")
	}

	return projects, err
}

func (s *service) checkIfProjectFound(ID string) (Project, error) {
	p, err := s.repository.FindByID(ID)
	if err != nil {
		return Project{}, errors.Wrap(err, fmt.Sprintf("The project %s not found", p.Name))
	}
	return p, nil
}
