package project_test

import (
	"errors"
	"testing"
	"time"

	"github.com/aperezg/feature-flags/project"
	store "github.com/aperezg/feature-flags/store/mock/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const projectID = "01C9S9Z0VG0MJFWYQZBFSSXDZE"
const projectName = "test_project"

func TestReturnAnErrorWhenProjectAlreadyExists(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(dummyProject(projectName), nil)

	s := project.NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())
}

func TestReturnAnErrorCreatingProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.CreateProject(projectName)
	assert.EqualError(t, err, err.Error())
}

func TestCreateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	p, err := s.CreateProject(projectName)
	assert.Equal(t, projectName, p.Name)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNameAlreadyExistsOnModifyProjectName(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(dummyProject(projectName), nil)

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnModifyProjectName(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistsOnModifyProjectName(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	_, err := s.ModifyProjectName(projectID, projectName)
	assert.Error(t, err)
}

func TestModifyProjectName(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByName", projectName).Return(project.Project{}, nil)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	p, err := s.ModifyProjectName(projectID, projectName)
	assert.Equal(t, projectName, p.Name)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnDeactivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistOnDeactivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.Error(t, err)
}

func TestDeactivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	err := s.DeactivateProject(projectID)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnActivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.Error(t, err)
}

func TestReturnAnErrorWhenPersistOnActivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(errors.New("error"))

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.Error(t, err)
}

func TestActivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Persist", mock.AnythingOfType("*project.Project")).Return(nil)

	s := project.NewService(r)
	err := s.ActivateProject(projectID)
	assert.NoError(t, err)
}

func TestReturnAnErrorWhenProjectNotFoundOnRemoveProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(project.Project{}, errors.New("error"))

	s := project.NewService(r)
	err := s.RemoveProject(projectID)
	assert.Error(t, err)

}

func TestReturnAnErrorWhenRemoveOnActivateProject(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Remove", projectID).Return(errors.New("error"))

	s := project.NewService(r)
	err := s.RemoveProject(projectID)
	assert.Error(t, err)
}

func TestReturnRemoveProject(t *testing.T) {
	r := new(store.RepositoryMock)

	r.On("FindByID", projectID).Return(dummyProject(projectName), nil)
	r.On("Remove", projectID).Return(nil)

	s := project.NewService(r)
	err := s.RemoveProject(projectID)
	assert.NoError(t, err)
}

func TestReturnAnErrorTryingToFetchResulsetOnFindProjects(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindAll").Return([]project.Project{}, errors.New("error"))

	s := project.NewService(r)
	projects, err := s.FindProjects()
	assert.Error(t, err)
	assert.Empty(t, projects)
}

func TestReturnAnEmptyResultWhenNotFoundProjectOnFindProjects(t *testing.T) {
	r := new(store.RepositoryMock)
	r.On("FindAll").Return([]project.Project{}, nil)

	s := project.NewService(r)
	projects, err := s.FindProjects()
	assert.NoError(t, err)
	assert.Empty(t, projects)
}

func TestReturnProjectsOnFindProjects(t *testing.T) {
	r := new(store.RepositoryMock)
	p := dummyProject(projectName)

	r.On("FindAll").Return([]project.Project{p}, nil)

	s := project.NewService(r)
	projects, err := s.FindProjects()
	assert.NoError(t, err)
	assert.NotEmpty(t, projects)
}

func dummyProject(name string) project.Project {
	return project.Project{
		ID:        projectID,
		Name:      name,
		CreatedAt: time.Now(),
		Status:    1,
	}
}
