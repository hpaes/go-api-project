package logger

type LogHandler interface {
	LogInformation(format string, args ...interface{})
	LogError(format string, args ...interface{})
}
