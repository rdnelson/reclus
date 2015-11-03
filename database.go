package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	SchemaTable = "VersionInfo"

	SchemaQuery = "SELECT SchemaVersion FROM VersionInfo"
	ProbeQuery  = "SELECT 1 FROM %s"
)

func setupDatabase(config *Config) (db *sql.DB, err error) {

	switch config.Database.Backend {
	case SQLite3:
		db, err = setupSqlite3DB(config)

		if err != nil {
			return
		}

	default:
		return nil, fmt.Errorf("Unknown backend '%s', failed to create DB.", config.Database.Backend)
	}

	if !probeTable(db, SchemaTable) {
		err = populateDatabase(db)

		if err != nil {
			db.Close()
			return nil, err
		}
	}

	_, err = checkSchema(db)

	if err != nil {
		db.Close()

		log.Fatal(err)
	}

	return db, err
}

func checkSchema(db *sql.DB) (int, error) {
	row := db.QueryRow(SchemaQuery)

	var version int

	if err := row.Scan(&version); err != nil {
		return 0, err
	}

	return version, nil
}

func probeTable(db *sql.DB, table string) bool {
	_, err := db.Exec(fmt.Sprintf(ProbeQuery, table))

	return err == nil
}

func populateDatabase(db *sql.DB) error {

	bytes, err := ioutil.ReadFile("sql/schema.sql")

	if err != nil {
		return err
	}

	str := string(bytes)
	sqlCmds := strings.Split(str, ";\r")

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, q := range sqlCmds {
		_, err := tx.Exec(q)

		if err != nil {
			return err
		}
	}

	err = tx.Commit()

	return err
}
