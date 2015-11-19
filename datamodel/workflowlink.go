package datamodel

type WorkflowLink struct {
	Id       int
	Workflow *Workflow
	Link     *StageLink
}
