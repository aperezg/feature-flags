package inmem

import "github.com/aperezg/feature-flags/project"

type ProjectInmem struct {
	projects map[string]project.Project
}
