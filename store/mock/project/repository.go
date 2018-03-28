package project

import (
	. "github.com/aperezg/feature-flags/identity"
	"github.com/aperezg/feature-flags/project"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock the struct fake for testing the Project callings to dbs
type RepositoryMock struct {
	mock.Mock
}

// FindByName function mock to simulate the calling on "FindByName"
func (m *RepositoryMock) FindByName(name string) (project.Project, error) {
	args := m.Called(name)
	return args.Get(0).(project.Project), args.Error(1)
}

// FindByID function mock to simulate the calling on "FindByID"
func (m *RepositoryMock) FindByID(ID Identity) (project.Project, error) {
	args := m.Called(ID)
	return args.Get(0).(project.Project), args.Error(1)
}

// Persist function mock to simulate the calling on "Persist"
func (m *RepositoryMock) Persist(project *project.Project) error {
	args := m.Called(project)
	return args.Error(0)
}
