package config

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/rdnelson/reclus/log"
)

type LoggingConfig struct {
	Level string
}

func (l LoggingConfig) Validate() error {

	l.Level = strings.ToUpper(l.Level)

	if l.Level == "" {
		l.Level = "INFO"
	}

	switch l.Level {
	case "DEBUG":
	case "INFO":
	case "WARN":
	case "ERROR":
		break
	default:
		log.Log.Warnf("Invalid logging level '%s', defaulting to INFO", l.Level)
		l.Level = "INFO"
	}

	return nil
}

func (l LoggingConfig) GetLogLevel() logrus.Level {
	switch l.Level {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
		break
	}

	return logrus.InfoLevel
}
