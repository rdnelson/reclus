package datamodel

type IssueType struct {
	Id       int
	Name     string
	Workflow *Workflow
}

type IssueTypeRepo interface {
	CreateIssueType(*IssueType) (string, error)
	GetIssueType(string) (*IssueType, error)
	UpdateIssueType(string, *IssueType) error
}
