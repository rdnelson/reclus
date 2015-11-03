package main

import (
	"database/sql"
	"errors"

	"gopkg.in/authboss.v0"
)

type User struct {
	ID   int
	Name string

	Email    string
	Password string
}

const (
	InsertQuery = "INSERT INTO Users (Key, Email, Password, Name) VALUES ($1, $2, $3, $4)"
	SelectQuery = "SELECT ID, Email, Password, Name FROM Users WHERE Key = $1"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (s UserRepo) Put(key string, attr authboss.Attributes) error {
	user := &User{}

	log.Debugf("Putting entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(user, false); err != nil {
		return err
	}

	_, err := s.db.Exec(InsertQuery, key, user.Email, user.Password, user.Name)

	return err
}

func (s UserRepo) Get(key string) (interface{}, error) {

	user := User{}

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
		return nil, authboss.ErrUserNotFound
	} else if userCount != 1 {
		return nil, errors.New("Invalid number of hits returned")
	}

	return &user, nil
}

func (s UserRepo) Create(key string, attr authboss.Attributes) error {
	var user User

	log.Debugf("Creating entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	_, err := s.db.Exec(InsertQuery, key, user.Email, user.Password, user.Name)

	return err
}
