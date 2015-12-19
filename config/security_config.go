package config

import (
	"encoding/base64"

	"github.com/rdnelson/reclus/log"
)

type SecurityConfig struct {
	SessionStoreKey string
	CookieStoreKey  string
}

func (s SecurityConfig) Validate() error {

	if s.SessionStoreKey != "" {
		key, err := base64.StdEncoding.DecodeString(s.SessionStoreKey)

		if err != nil {
			return err
		}

		if len(key) != 64 {
			log.Log.Warnf("Session Store Key is not 64 bytes long, it's %d bytes.", len(key))
		}
	}

	if s.CookieStoreKey != "" {
		key, err := base64.StdEncoding.DecodeString(s.CookieStoreKey)

		if err != nil {
			return err
		}

		if len(key) != 64 {
			log.Log.Warnf("Cookie Store Key is not 64 bytes long, it's %d bytes.", len(key))
		}
	}

	return nil
}
