package main

import (
	"gopkg.in/authboss.v0"
)

type User struct {
	ID   int
	Name string

	Email    string
	Password string
}

type UserRepo struct {
	db Database
}

func NewUserRepo(db Database) *UserRepo {
	return &UserRepo{db}
}

func (s UserRepo) Put(key string, attr authboss.Attributes) error {
	user := &User{}

	log.Debugf("Putting entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(user, false); err != nil {
		return err
	}

	return s.db.UpdateUser(key, user)
}

func (s UserRepo) Get(key string) (interface{}, error) {
	log.Debugf("Getting entry '%s'", key)

	user, err := s.db.GetUser(key)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, authboss.ErrUserNotFound
	}

	return user, nil
}

func (s UserRepo) Create(key string, attr authboss.Attributes) error {
	var user User

	log.Debugf("Creating entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	return s.db.AddUser(key, &user)
}
