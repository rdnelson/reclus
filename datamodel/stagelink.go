package datamodel

type StageLink struct {
	Id          int
	Name        string
	Description string
	FromStage   *WorkflowStage
	ToStage     *WorkflowStage
}
