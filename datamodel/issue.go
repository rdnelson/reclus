package datamodel

import (
	"time"
)

type Issue struct {
	// Issue info
	Id    int
	Title string
	Type  *IssueType

	// Users
	AssignedTo *User
	Reporter   *User

	// Dates
	CreatedOn   time.Time
	LastUpdated time.Time
}

type IssueRepo interface {
	CreateIssue(*Issue) (string, error)
	GetIssue(string) (*Issue, error)
	UpdateIssue(string, *Issue) error
}
