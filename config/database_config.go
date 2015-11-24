package config

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	Backend string
}

var (
	SupportedBackends = make(map[string]func(interface{}) error)
)

func RegisterDbBackend(name string, config interface{}, validate func(interface{}) error) {
	name = strings.ToLower(name)
	SupportedBackends[name] = validate

	if Cfg.Backend == nil {
		Cfg.Backend = make(map[string]interface{})
	}

	Cfg.Backend[name] = config
}

func (d *DBConfig) Validate() error {
	backendName := strings.ToLower(d.Backend)

	validate, found := SupportedBackends[backendName]

	if !found {
		return fmt.Errorf("Invalid database backend: '%s'\n", backendName)
	}

	validate(Cfg.Backend[backendName])

	// Normalize the string
	d.Backend = backendName

	return nil
}
