package main

import (
	"encoding/base64"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"gopkg.in/authboss.v0"
)

const SessionCookieName = "reclus_session"

var (
	sessionStore *sessions.CookieStore
)

type SessionStore struct {
	writer  http.ResponseWriter
	request *http.Request
}

func InitializeSessionStore(config *Config) error {
	if config.Security.SessionStoreKey == "" {
		config.Security.SessionStoreKey = base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(64))
		log.Debugf("Generated Session Store Key: '%s'", config.Security.SessionStoreKey)
	}

	key, err := base64.StdEncoding.DecodeString(config.Security.SessionStoreKey)

	if err != nil {
		return err
	}

	sessionStore = sessions.NewCookieStore(key)

	return nil
}

func NewSessionStore(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStore{w, r}
}

func (s SessionStore) Get(key string) (string, bool) {
	session, err := sessionStore.Get(s.request, SessionCookieName)

	if err != nil {
		return "", false
	}

	strInf, ok := session.Values[key]

	if !ok {
		log.Debugf("Failed to find data in sessionStore for key '%s'", key)
		return "", false
	}

	str, ok := strInf.(string)

	return str, ok
}

func (s SessionStore) Put(key, value string) {
	session, err := sessionStore.Get(s.request, SessionCookieName)

	if err != nil {
		return
	}

	log.Debugf("Saving data to sessionStore for key '%s'='%s'", key, value)
	session.Values[key] = value
	session.Save(s.request, s.writer)
}

func (s SessionStore) Del(key string) {
	session, err := sessionStore.Get(s.request, SessionCookieName)

	if err != nil {
		return
	}

	delete(session.Values, key)
	session.Save(s.request, s.writer)
}
