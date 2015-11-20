package datamodel

type StageLink struct {
	Id          int
	Name        string
	Description string
	FromStage   *WorkflowStage
	ToStage     *WorkflowStage
}

type StageLinkRepo interface {
	CreateStageLink(*StageLink) (string, error)
	GetStageLink(string) (*StageLink, error)
	UpdateStageLink(string, *StageLink) error
}
