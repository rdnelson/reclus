package config

import "fmt"

type ServerConfig struct {
	Hostname string
	Port     int
}

func (s ServerConfig) Validate() error {

	if s.Port <= 0 && s.Port >= 65536 {
		return fmt.Errorf("Invalid port '%d'. Must be between 1 and 65535 (inclusive.)", s.Port)
	}

	return nil
}
