package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rdnelson/reclus/datamodel"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SchemaTable = "VersionInfo"

	SchemaQuery = "SELECT SchemaVersion FROM VersionInfo"
	ProbeQuery  = "SELECT 1 FROM %s"

	UpdateQuery = "UPDATE Users SET (Email, Password, Name) VALUES ($2, $3, $4) WHERE Key = $1"
	InsertQuery = "INSERT INTO Users (Key, Email, Password, Name) VALUES ($1, $2, $3, $4)"
	SelectQuery = "SELECT ID, Email, Password, Name FROM Users WHERE Key = $1"
)

const (
	SQLite3 = "sqlite3"
)

type SQLite3Database struct {
	path string
	db   *sql.DB
}

type SQLiteConfig struct {
	Path string
}

func init() {
	SupportedBackends = append(SupportedBackends, SQLite3)
	DatabaseProviders[SQLite3] = NewSqlite3Db
}

func NewSqlite3Db(config *Config) (Database, error) {
	if err := config.SQLite3.Validate(); err != nil {
		return nil, err
	}

	return &SQLite3Database{config.SQLite3.Path, nil}, nil
}

func (s *SQLiteConfig) Validate() error {
	path, err := filepath.Abs(s.Path)

	if err != nil {
		return fmt.Errorf("Invalid path to *SQLite database: '%v'\n", err)
	}

	s.Path = path

	return nil
}

func (db *SQLite3Database) Open() (err error) {
	log.Debugf("Opening database: '%s'", db.path)

	if _, err := os.Stat(db.path); os.IsNotExist(err) {
		err = db.Create()

		if err != nil {
			return err
		}
	}

	db.db, err = sql.Open(SQLite3, db.path)

	if err != nil {
		return err
	}

	return db.db.Ping()
}

func (db *SQLite3Database) Create() error {
	log.Debugf("Creating sqlite file: '%s'", db.path)

	dbPath := db.path
	dbDir := filepath.Dir(dbPath)

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (db *SQLite3Database) ValidateSchema() error {
	version := -1

	row := db.db.QueryRow(SchemaQuery)

	if err := row.Scan(&version); err != nil {
		return err
	}

	if version < 0 {
		return fmt.Errorf("Invalid schema version: %d", version)
	}

	return nil
}

func (db *SQLite3Database) Close() error {
	return db.db.Close()
}

func (db *SQLite3Database) PopulateSchema() error {
	bytes, err := ioutil.ReadFile("sql/schema.sql")

	if err != nil {
		return err
	}

	str := string(bytes)
	sqlCmds := strings.Split(str, ";\r")

	tx, err := db.db.Begin()

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

func (db *SQLite3Database) probeTable(table string) bool {
	_, err := db.db.Exec(fmt.Sprintf(ProbeQuery, table))

	return err == nil
}

func (s *SQLite3Database) UpdateUser(key string, user *datamodel.User) error {
	_, err := s.db.Exec(UpdateQuery, key, user.Email, user.Password, user.Name)

	return err
}

func (s *SQLite3Database) GetUser(key string) (*datamodel.User, error) {
	user := datamodel.User{}

	log.Debugf("Getting entry '%s'", key)
	rows, err := s.db.Query(SelectQuery, key)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userCount := 0

	for rows.Next() {
		userCount++
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, err
		}
	}

	log.Debugf("Found '%d' matching user entries", userCount)

	if userCount == 0 {
		return nil, nil
	} else if userCount != 1 {
		return nil, errors.New("Invalid number of hits returned")
	}

	return &user, nil
}

func (s *SQLite3Database) AddUser(key string, user *datamodel.User) error {
	_, err := s.db.Exec(InsertQuery, key, user.Email, user.Password, user.Name)

	return err
}
