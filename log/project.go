package log

import ()

type Project struct {
	Id   string
	Name string
}

func (p Project) String() string {
	return p.Name
}
