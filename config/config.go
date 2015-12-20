package config

import (
	"errors"
	"os"

	"gopkg.in/gcfg.v1"
)

const (
	ConfigPath = "cfg/reclus.cfg"
)

var (
	Cfg Config
)

type Config struct {
	Logging  LoggingConfig
	Server   ServerConfig
	Database DBConfig
	Security SecurityConfig
	Backend  map[string]interface{}
}

func Load() error {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		return errors.New("No configuration file was found.\n")
	}

	if err := gcfg.ReadFileInto(&Cfg, ConfigPath); err != nil {
		return err
	}

	return validateConfig()
}

func validateConfig() error {
	if err := Cfg.Server.Validate(); err != nil {
		return err
	}

	if err := Cfg.Database.Validate(); err != nil {
		return err
	}

	if err := Cfg.Security.Validate(); err != nil {
		return err
	}

	if err := Cfg.Logging.Validate(); err != nil {
		return err
	}

	return nil
}
