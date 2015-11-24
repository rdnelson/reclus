package datamodel

type VersionInfo struct {
	VersionNumber int
	IsActive      int
}

type VersionInfoRepo interface {
	CreateVersionInfo(*VersionInfo) (string, error)
	GetVersionInfo(string) (*VersionInfo, error)
	UpdateVersionInfo(string, *VersionInfo) error
}
