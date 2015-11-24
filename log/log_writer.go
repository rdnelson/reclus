package log

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogWriter struct {
	Logger *logrus.Logger
}

func (l LogWriter) Write(data []byte) (int, error) {
	l.Logger.Info(strings.TrimSpace(string(data)))

	return len(data), nil
}
