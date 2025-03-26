package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type LogLevel string

const (
	INFO  LogLevel = "info"
	ERROR LogLevel = "error"
	WARN  LogLevel = "warn"
)

var (
	_ BaseLogger = (*Logrus)(nil)
)

// Logrus BaseLogger implementation based on logrus.
type Logrus struct {
}

// NewLogrus returns def logger
func NewLogrus(levelRaw string) (BaseLogger, error) {
	level, err := logrus.ParseLevel(levelRaw)
	if err != nil {
		return nil, fmt.Errorf("parse Level from config: %w", err)
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	return &Logrus{}, nil
}

func (l *Logrus) Infof(msg string, fields ...interface{}) {
	logrus.Infof(msg, fields...)
}

func (l *Logrus) Warnf(msg string, fields ...interface{}) {
	logrus.Warnf(msg, fields...)
}

func (l *Logrus) Errorf(msg string, fields ...interface{}) {
	logrus.Errorf(msg, fields...)
}

func (l *Logrus) Fatalf(msg string, fields ...interface{}) {
	logrus.Errorf(msg, fields...)
}
