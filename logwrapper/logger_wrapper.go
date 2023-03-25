package logwrapper

import (
  "github.com/sirupsen/logrus"
)

type Event struct {
  id      int
  message string
}

type Logger struct {
  *logrus.Logger
}

func NewLogger() *Logger {
	var baseLogger = logrus.New()
	var logger = &Logger{baseLogger}
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}

var (
  databaseConnectionFailureMessage      = Event{1, "Failed to connect to database: %s"}
  serverStartFailureMessage = Event{2, "Failed to start server on port: %d"}
)

func (l *Logger) DatabaseConnectionFailure(argumentName string) {
  l.Panicf(databaseConnectionFailureMessage.message, argumentName)
}

func (l *Logger) ServerStartFailure(argumentName int) {
  l.Errorf(serverStartFailureMessage.message, argumentName)
}