package backends

import (
	"fmt"

	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/datamodel"
)

var DatabaseProviders = make(map[string]func() (Database, error))

type Database interface {
	// Database management functions
	Open() error
	MigrateSchema() error
	Close() error

	// User Management Functions
	datamodel.FullRepo
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

	if err = db.MigrateSchema(); err != nil {
		return nil, err
	}

	return db, nil
}
