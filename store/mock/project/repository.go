package project

import (
	"github.com/aperezg/feature-flags/project"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock the struct fake for testing the Project callings to dbs
type RepositoryMock struct {
	mock.Mock
}

// CreateProject function mock to simulate the calling on "CreateProject"
func (m *RepositoryMock) CreateProject(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

// FindByName function mock to simulate the calling on "FindByName"
func (m *RepositoryMock) FindByName(name string, project *project.Project) error {
	args := m.Called(name, project)
	return args.Error(0)
}
