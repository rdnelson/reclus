package main

import (
	"github.com/rdnelson/reclus/backends"
	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"

	"gopkg.in/authboss.v0"
)

type AuthUserRepo struct {
	db backends.Database
}

func NewUserRepo(db backends.Database) *AuthUserRepo {
	return &AuthUserRepo{db}
}

func (s AuthUserRepo) Put(key string, attr authboss.Attributes) error {
	user := &datamodel.User{Email: key}

	log.Log.Debugf("Putting entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(user, false); err != nil {
		return err
	}

	return s.db.UpdateUser(user)
}

func (s AuthUserRepo) Get(key string) (interface{}, error) {
	log.Log.Debugf("Getting partial user '%s'", key)

	user, err := s.db.GetPartialUser(key)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, authboss.ErrUserNotFound
	}

	return user, nil
}

func (s AuthUserRepo) GetUser(user *datamodel.User) (*datamodel.User, error) {
	log.Log.Debugf("Getting user '%s'", user.Email)

	user, err := s.db.GetUser(user.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, authboss.ErrUserNotFound
	}

	return user, nil
}

func (s AuthUserRepo) Create(key string, attr authboss.Attributes) error {
	var user datamodel.User
	user.Email = key

	log.Log.Debugf("Creating entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	return s.db.AddUser(&user)
}
