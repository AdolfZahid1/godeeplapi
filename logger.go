package godeeplapi

import (
	"log"
	"os"
)

// Logger defines the interface for logging
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// defaultLogger is a simple implementation using the standard log package
type defaultLogger struct {
	logger *log.Logger
}

// NewDefaultLogger Initialize the logger field with default values
func NewDefaultLogger() *defaultLogger {
	return &defaultLogger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *defaultLogger) Debug(msg string, args ...interface{}) {
	if l.logger == nil {
		l.logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	l.logger.Printf("[DEBUG] "+msg, args...)
}

func (l *defaultLogger) Info(msg string, args ...interface{}) {
	if l.logger == nil {
		l.logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	l.logger.Printf("[INFO] "+msg, args...)
}

func (l *defaultLogger) Error(msg string, args ...interface{}) {
	if l.logger == nil {
		l.logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	l.logger.Printf("[ERROR] "+msg, args...)
}
