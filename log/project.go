package log

import (
	"github.com/johnweldon/go/util"
)

type Project struct {
	id   util.UUID
	Name string
}

func (p Project) String() string {
	return p.Name
}

func NewProject(name string) Project {
	return Project{util.NewUUID(), name}
}
