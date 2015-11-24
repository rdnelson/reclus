package backends

import (
	"fmt"

	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"
)

var DatabaseProviders = make(map[string]func() (Database, error))

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

func NewDatabase() (db Database, err error) {
	factory := DatabaseProviders[config.Cfg.Database.Backend]

	if factory == nil {
		return nil, fmt.Errorf("Unknown database provider: %v", config.Cfg.Database.Backend)
	}

	db, err = factory()

	if err != nil {
		return nil, err
	}

	if err = db.Open(); err != nil {
		return nil, err
	}

	log.Log.Debugf("Database: '%v'", db)

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
