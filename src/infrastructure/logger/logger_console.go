package logger

import "log"

type ConsoleLogger struct{}

func (c *ConsoleLogger) LogInformation(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (c *ConsoleLogger) LogError(format string, args ...interface{}) {
	log.Printf(format, args...)
}
