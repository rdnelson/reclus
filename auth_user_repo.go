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
	user := &datamodel.User{}

	log.Log.Debugf("Putting entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(user, false); err != nil {
		return err
	}

	return s.db.UpdateUser(key, user)
}

func (s AuthUserRepo) Get(key string) (interface{}, error) {
	log.Log.Debugf("Getting entry '%s'", key)

	user, err := s.db.GetUser(key)

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

	log.Log.Debugf("Creating entry '%s' with attributes: '%v'", key, attr)
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	return s.db.AddUser(key, &user)
}
