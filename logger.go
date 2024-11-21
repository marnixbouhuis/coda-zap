package codazap

import (
	"github.com/marnixbouhuis/coda"
	"go.uber.org/zap"
)

type Logger struct {
	l *zap.Logger
}

var _ coda.Logger = &Logger{}

// NewLogger creates a new coda logger that sends all log messages to zap.

func NewLogger(l *zap.Logger) *Logger {
	return &Logger{l: l}
}

func (c *Logger) Info(str string) {
	c.l.Info(str)
}

func (c *Logger) Error(str string) {
	c.l.Error(str)
}
