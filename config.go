package main

import (
	"errors"
	"os"

	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
)

const (
	ConfigPath = "cfg/reclus.cfg"
)

type Config struct {
	Database DBConfig
	Security SecurityConfig
	SQLite3  SQLiteConfig
}

func loadConfig(config *Config) error {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		return errors.New("No configuration file was found.\n")
	}

	sconf.Must(config).Read(ini.File(ConfigPath))

	return validateConfig(config)
}

func validateConfig(config *Config) error {
	if err := validateDatabase(config); err != nil {
		return err
	}

	return nil
}
