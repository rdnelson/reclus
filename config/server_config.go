package config

import "fmt"

type ServerConfig struct {
	Hostname    string
	Port        int
	DisplayName string
}

func (s ServerConfig) Validate() error {

	if s.Port <= 0 && s.Port >= 65536 {
		return fmt.Errorf("Invalid port '%d'. Must be between 1 and 65535 (inclusive.)", s.Port)
	}

	host := s.Hostname

	if host == "" {
		host = "localhost"
	}

	if s.DisplayName == "" {
		s.DisplayName = fmt.Sprintf("http://%s:%d", host, s.Port)
	}

	return nil
}
