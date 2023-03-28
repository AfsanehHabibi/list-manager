package logwrapper

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Event struct {
	id      int
	message string
}

type Logrus struct {
	*logrus.Logger
}

func NewLogger() *Logrus {
	var baseLogger = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		baseLogger.SetOutput(file)
	} else {
		baseLogger.Info("Failed to log to file, using default stderr")
	}
	var logger = &Logrus{baseLogger}
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}

var (
	databaseConnectionFailureMessage = Event{1, "Failed to connect to database"}
	serverStartFailureMessage        = Event{2, "Failed to start server"}
)

func (l *Logrus) DatabaseConnectionFailure(databaseName string, err error) {
	l.WithFields(logrus.Fields{
		"error": err,
		"id":    databaseConnectionFailureMessage.id,
	}).Panic(databaseConnectionFailureMessage.message)
}

func (l *Logrus) ServerStartFailure(port int, err error) {
	l.WithFields(logrus.Fields{
		"error": err,
		"id":    serverStartFailureMessage.id,
	}).Error(serverStartFailureMessage.message)
}
