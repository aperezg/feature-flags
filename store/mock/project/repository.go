package project

import (
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
