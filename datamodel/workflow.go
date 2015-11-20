package datamodel

type Workflow struct {
	Id         int
	Name       string
	FirstStage *WorkflowStage
}

type WorkflowRepo interface {
	CreateWorkflow(*Workflow) (string, error)
	GetWorkflow(string) (*Workflow, error)
	UpdateWorkflow(string, *Workflow) error
}
