package inmem

import (
	"sync"

	"github.com/aperezg/feature-flags/project"
)

type projectRepository struct {
	mtx      sync.RWMutex
	projects map[string]*project.Project
}

// NewProjectRepository returns a new instance of a in-memory project repository.
func NewProjectRepository() project.Repository {
	return &projectRepository{
		projects: make(map[string]*project.Project),
	}
}

func (r *projectRepository) FindAll() ([]*project.Project, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	p := make([]*project.Project, 0, len(r.projects))
	for _, val := range r.projects {
		p = append(p, val)
	}
	return p, nil
}
func (r *projectRepository) FindByName(name string) (project.Project, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for _, val := range r.projects {
		if val.Name == name {
			return *val, nil
		}
	}
	return project.Project{}, nil
}
func (r *projectRepository) FindByID(ID string) (project.Project, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return *r.projects[ID], nil
}
func (r *projectRepository) Persist(project *project.Project) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.projects[project.ID] = project
	return nil

}
func (r *projectRepository) Remove(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.projects, ID)
	return nil
}
