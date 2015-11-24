package log

import (
	"github.com/sirupsen/logrus"
)

var (
	Log = logrus.New()
)

const (
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
)
