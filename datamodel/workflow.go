package datamodel

type Workflow struct {
	Id         int
	Name       string
	FirstStage *WorkflowStage
}
