package project_test

import (
	"errors"
	. "github.com/aperezg/feature-flags/identity"
	"github.com/aperezg/feature-flags/project"
	storeProject "github.com/aperezg/feature-flags/store/mock/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

const projectID = Identity("f30c7ee9-ef5a-4c7f-b680-6f1166f3451b")
const projectName = "test_project"

func TestReturnAnErrorWhenProjectAlreadyExists(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(dummyProject(projectName), nil)

	s := project.NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())
}

func TestReturnAnErrorCreatingProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.CreateProject(projectName)
	assert.EqualError(t, err, err.Error())
}

func TestCreateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	p, err := s.CreateProject(projectName)
	assert.Equal(t, projectName, p.Name)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNameAlreadyExistsOnModifyProjectName(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(dummyProject(projectName), nil)

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnModifyProjectName(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistsOnModifyProjectName(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestModifyProjectName(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	p, err := s.ModifyProjectName(projectID, projectName)
	assert.Equal(t, projectName, p.Name)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnDeactivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistOnDeactivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.Error(t, err)
}

func TestDeactivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnActivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistOnActivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.Error(t, err)
}

func TestActivateProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnRemoveProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.RemoveProject(projectID)
	assert.Error(t, err)

}

func dummyProject(name string) project.Project {
	return project.Project{
		ID:        projectID,
		Name:      name,
		CreatedAt: time.Now(),
		Status:    1,
	}
}
