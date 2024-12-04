package logger

import "log"

type (
	LogHandler interface {
		LogInformation(format string, args ...interface{})
		LogError(format string, args ...interface{})
	}
	logHandler struct {
	}
)

func NewConsoleLogger() LogHandler {
	return &logHandler{}
}

func (l *logHandler) LogInformation(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *logHandler) LogError(format string, args ...interface{}) {
	log.Printf(format, args...)
}
