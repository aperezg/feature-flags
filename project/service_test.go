package project

import (
	"errors"
	storeProject "github.com/aperezg/feature-flags/store/mock/project"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnAnErrorCreatingProject(t *testing.T) {
	r := new(storeProject.RepositoryMock)
	r.On("CreateProject", "test_project").Return(errors.New("error"))

	s := NewService(r)
	_, err := s.CreateProject("test_project")
	assert.EqualError(t, err, err.Error())

}
