package main

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"

	"gopkg.in/authboss.v0"
)

var cookieStore *securecookie.SecureCookie

type CookieStore struct {
	writer  http.ResponseWriter
	request *http.Request
}

func InitializeCookieStore(config *Config) error {
	if config.Security.CookieStoreKey == "" {
		config.Security.CookieStoreKey = base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(64))
	}

	key, err := base64.StdEncoding.DecodeString(config.Security.CookieStoreKey)

	if err != nil {
		return err
	}

	cookieStore = securecookie.New(key, nil)

	return nil
}

func NewCookieStore(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStore{w, r}
}

func (c CookieStore) Get(key string) (string, bool) {
	cookie, err := c.request.Cookie(key)

	if err != nil {
		return "", false
	}

	var value string
	if err = cookieStore.Decode(key, cookie.Value, &value); err != nil {
		return "", false
	}

	return value, true
}

func (c CookieStore) Put(key, value string) {
	encoded, err := cookieStore.Encode(key, value)

	if err != nil {
		return
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}

	http.SetCookie(c.writer, cookie)
}

func (c CookieStore) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}

	http.SetCookie(c.writer, cookie)
}
