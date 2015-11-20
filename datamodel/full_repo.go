package datamodel

type FullRepo interface {
	IssueRepo
	IssueTypeRepo
	ProjectRepo
	ProjectIssueTypeRepo
	StageLinkRepo
	UserRepo
	WorkflowRepo
	WorkflowLinkRepo
	WorkflowStageRepo
}
