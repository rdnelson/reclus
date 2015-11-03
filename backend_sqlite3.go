package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func validateSQLite3(config *Config) error {
	path, err := filepath.Abs(config.SQLite3.Path)

	if err != nil {
		return fmt.Errorf("Invalid path to SQLite database: '%v'\n", err)
	}

	config.SQLite3.Path = path

	return nil
}

func setupSqlite3DB(config *Config) (db *sql.DB, err error) {
	dbPath := config.SQLite3.Path

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		db, err = createSqlite3DB(config)
	} else {
		db, err = sql.Open(SQLite3, dbPath)
	}

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return
}

func createSqlite3DB(config *Config) (db *sql.DB, err error) {
	dbPath := config.SQLite3.Path
	dbDir := filepath.Dir(dbPath)

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return nil, err
		}
	}

	return sql.Open(SQLite3, dbPath)
}
