package datamodel

type ProjectIssueType struct {
	Id        int
	Project   *Project
	IssueType *IssueType
}

type ProjectIssueTypeRepo interface {
	CreateProjectIssueType(*ProjectIssueType) (string, error)
	GetProjectIssueType(string) (*ProjectIssueType, error)
	UpdateProjectIssueType(string, *ProjectIssueType) error
}
