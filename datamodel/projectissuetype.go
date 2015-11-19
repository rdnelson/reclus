package datamodel

type ProjectIssueType struct {
	Id        int
	Project   *Project
	IssueType *IssueType
}
