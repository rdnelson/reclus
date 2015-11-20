package main

import (
	"fmt"

	"github.com/rdnelson/reclus/datamodel"
)

var DatabaseProviders = make(map[string]func(*Config) (Database, error))

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

	factory := DatabaseProviders[config.Database.Backend]

	if factory == nil {
		return nil, fmt.Errorf("Unknown database provider: %v", config.Database.Backend)
	}

	db, err = factory(config)

	if err != nil {
		return nil, err
	}

	if err = db.Open(); err != nil {
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
