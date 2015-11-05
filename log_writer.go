package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogWriter struct {
	logger *logrus.Logger
}

func (l LogWriter) Write(data []byte) (int, error) {
	l.logger.Info(strings.TrimSpace(string(data)))

	return len(data), nil
}
