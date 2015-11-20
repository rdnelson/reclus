package datamodel

type WorkflowLink struct {
	Id       int
	Workflow *Workflow
	Link     *StageLink
}

type WorkflowLinkRepo interface {
	CreateWorkflowLink(*WorkflowLink) (string, error)
	GetWorkflowLink(string) (*WorkflowLink, error)
	UpdateWorkflowLink(string, *WorkflowLink) error
}
