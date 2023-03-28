package logwrapper

type Logger interface {
	NewLogger() Logger
	DatabaseConnectionFailure(databaseName string, err error)
	ServerStartFailure(port int, err error)
}