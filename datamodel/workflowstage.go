package datamodel

type WorkflowStage struct {
	Id   int
	Name string
}

type WorkflowStageRepo interface {
	CreateWorkflowStage(*WorkflowStage) (string, error)
	GetWorkflowStage(string) (*WorkflowStage, error)
	UpdateWorkflowStage(string, *WorkflowStage) error
}
