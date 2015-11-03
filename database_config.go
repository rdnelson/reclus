package main

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	Backend string
}

type SQLiteConfig struct {
	Path string
}

const (
	SQLite3 = "sqlite3"
)

var (
	SupportedBackends = [...]string{SQLite3}
)

func validateDatabase(config *Config) error {
	backend := strings.ToLower(config.Database.Backend)

	if _, err := ListContains(SupportedBackends[:], backend); err != nil {
		return fmt.Errorf("Invalid database backend: '%s'\n", backend)
	}

	// Normalize the string
	config.Database.Backend = backend

	switch backend {
	case SQLite3:
		return validateSQLite3(config)

	default:
		return fmt.Errorf("Unexpected database backend: '%s'\n", config.Database.Backend)
	}

	return nil
}
