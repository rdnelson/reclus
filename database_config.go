package main

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	Backend string
}

const (
	SQLite3 = "sqlite3"
)

var (
	SupportedBackends = [...]string{SQLite3}
)

func (d *DBConfig) Validate() error {
	backendName := strings.ToLower(d.Backend)

	if _, err := ListContains(SupportedBackends[:], backendName); err != nil {
		return fmt.Errorf("Invalid database backend: '%s'\n", backendName)
	}

	// Normalize the string
	d.Backend = backendName

	return nil
}
