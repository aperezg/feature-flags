package project_test

import (
	"errors"
	"github.com/aperezg/feature-flags/project"
	storeProject "github.com/aperezg/feature-flags/store/mock/project"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnAnErrorCreatingProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("CreateProject", "test_project").Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())
}

func TestErrorAfterCreateProject(t *testing.T) {
	var p project.Project
	projectName := "test_project"

	r := new(storeProject.RepositoryMock)
	r.On("CreateProject", projectName).Return(nil)
	r.On("FindByName", projectName, &p).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())
}

func TestCreateProject(t *testing.T) {
	var p project.Project
	projectName := "test_project"

	r := new(storeProject.RepositoryMock)
	r.On("CreateProject", projectName).Return(nil)
	r.On("FindByName", projectName, &p).Return(nil)

	s := project.NewService(r)
	p, err := s.CreateProject("test_project")
	assert.Equal(t, project.Project{}, p)
	assert.NoError(t, err)
}
