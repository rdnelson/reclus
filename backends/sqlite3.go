package backends

import (
	"fmt"
	"path/filepath"

	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SQLite3 = "sqlite3"
)

type SQLite3Database struct {
	path string
	db   gorm.DB
}

type SQLiteConfig struct {
	Path string
}

func init() {
	config.RegisterDbBackend(SQLite3, &SQLiteConfig{}, func(cfg interface{}) error {
		return cfg.(*SQLiteConfig).Validate()
	})

	DatabaseProviders[SQLite3] = NewSqlite3Db
}

func NewSqlite3Db() (Database, error) {
	return &SQLite3Database{config.Cfg.Backend[SQLite3].(*SQLiteConfig).Path, gorm.DB{}}, nil
}

func (s *SQLiteConfig) Validate() error {
	path, err := filepath.Abs(s.Path)

	if err != nil {
		return fmt.Errorf("Invalid path to *SQLite database: '%v'\n", err)
	}

	s.Path = path

	return nil
}

func (db *SQLite3Database) MigrateSchema() error {
	db.db.AutoMigrate(&datamodel.User{})

	return nil
}

func (db *SQLite3Database) Open() (err error) {
	db.db, err = gorm.Open(SQLite3, db.path)

	return err
}

func (db *SQLite3Database) Close() error {
	return db.db.Close()
}

func (s *SQLite3Database) UpdateUser(user *datamodel.User) error {
	s.db.Save(user)

	return nil
}

func (s *SQLite3Database) GetUser(username string) (*datamodel.User, error) {
	usr := &datamodel.User{Email: username}

	if err := s.db.Where(usr).First(usr).Error; err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *SQLite3Database) GetPartialUser(username string) (*datamodel.User, error) {
	usr := &datamodel.User{Email: username}

	if err := s.db.Where(usr).Select("name, password").First(usr).Error; err != nil {
		log.Log.Warnf("Error getting partial user: %v", err)

		if err == gorm.RecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}

	}

	return usr, nil
}

func (s *SQLite3Database) AddUser(user *datamodel.User) error {
	s.db.Save(user)
	return nil
}
