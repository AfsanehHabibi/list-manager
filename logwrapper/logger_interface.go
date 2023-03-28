package logwrapper

type Logger interface {
	DatabaseConnectionFailure(databaseName string, err error)
	ServerStartFailure(port int, err error)
	ServerStartSuccess(port int)
	DatabaseConnectionSuccess(databaseName string)
}