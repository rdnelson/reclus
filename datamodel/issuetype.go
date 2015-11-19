package datamodel

type IssueType struct {
	Id       int
	Name     string
	Workflow *Workflow
}
