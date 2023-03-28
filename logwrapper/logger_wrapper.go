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

func NewLogger() Logger {
	var baseLogger = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		baseLogger.SetOutput(file)
	} else {
		baseLogger.Info("Failed to log to file, using default stderr")
	}
	var logger = Logrus{baseLogger}
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}

var (
	databaseConnectionFailureMessage = Event{1, "Failed to connect to database"}
	serverStartFailureMessage        = Event{2, "Failed to start server"}
	serverStartSuccessMessage        = Event{3, "Started server"}
	databaseConnectionSuccessMessage = Event{4, "Started a database connection"}
)

func (l Logrus) DatabaseConnectionFailure(databaseName string, err error) {
	l.WithFields(logrus.Fields{
		"error": err,
		"id":    databaseConnectionFailureMessage.id,
		"name":  databaseName,
	}).Panic(databaseConnectionFailureMessage.message)
}

func (l Logrus) ServerStartFailure(port int, err error) {
	l.WithFields(logrus.Fields{
		"error": err,
		"id":    serverStartFailureMessage.id,
		"port":  port,
	}).Error(serverStartFailureMessage.message)
}

func (l Logrus) ServerStartSuccess(port int) {
	l.WithFields(logrus.Fields{
		"id":   serverStartSuccessMessage.id,
		"port": port,
	}).Error(serverStartSuccessMessage.message)
}

func (l Logrus) DatabaseConnectionSuccess(databaseName string) {
	l.WithFields(logrus.Fields{
		"id":   databaseConnectionSuccessMessage.id,
		"name": databaseName,
	}).Info(databaseConnectionSuccessMessage.message)
}
