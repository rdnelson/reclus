package datamodel

type Project struct {
	Id          int
	Name        string
	Description string
}

type ProjectRepo interface {
	CreateProject(*Project) (string, error)
	GetProject(string) (*Project, error)
	UpdateProject(string, *Project) error
}
