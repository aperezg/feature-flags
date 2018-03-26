package project

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) CreateProject(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
