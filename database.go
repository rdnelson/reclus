package main

import (
	"fmt"

	"github.com/rdnelson/reclus/datamodel"
)

type Database interface {
	// Database management functions
	Open() error
	Create() error
	ValidateSchema() error
	PopulateSchema() error
	Close() error

	// User Management Functions
	datamodel.UserRepo
}

func NewDatabase(config *Config) (db Database, err error) {
	switch config.Database.Backend {
	case SQLite3:

		if err := config.SQLite3.Validate(); err != nil {
			return nil, err
		}

		db = &SQLite3Database{config.SQLite3.Path, nil}
		break

	default:
		return nil, fmt.Errorf("Unexpected database backend: '%s'\n", config.Database.Backend)
	}

	if err := db.Open(); err != nil {
		return nil, err
	}

	log.Debugf("Database: '%v'", db)

	if err := db.ValidateSchema(); err != nil {
		if err := db.PopulateSchema(); err != nil {
			db.Close()
			return nil, err
		}
	}

	if err := db.ValidateSchema(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
