package project_test

import (
	"errors"
	"github.com/aperezg/feature-flags/project"
	storeProject "github.com/aperezg/feature-flags/store/mock/project"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestReturnAnErrorWhenProjectAlreadyExists(t *testing.T) {
	projectName := "test_project"

	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(dummyProject(projectName), nil)

	s := project.NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())
}

func TestReturnAnErrorCreatingProject(t *testing.T) {
	projectName := "test_project"

	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.CreateProject(projectName)
	assert.EqualError(t, err, err.Error())
}

func TestCreateProject(t *testing.T) {
	projectName := "test_project"

	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	p, err := s.CreateProject("test_project")
	assert.Equal(t, projectName, p.Name)
	assert.NoError(t, err)
}

func dummyProject(name string) project.Project {
	return project.Project{
		ID:        uuid.NewV4().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Status:    1,
	}
}
