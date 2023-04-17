package main

import (
	"os"
	"github.com/joho/godotenv"
	"database/sql"
	"github.com/AfsanehHabibi/list-manager/handler"
	"github.com/AfsanehHabibi/list-manager/logwrapper"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	logger := logwrapper.NewLogger()
	e := echo.New()
	h := handler.NewHandler("list")
	envErr := godotenv.Load()
  	if envErr != nil {
    	logger.DatabaseConnectionFailure("psql", envErr)
  	}
	connStr := os.Getenv("DB_STR_CONNECTION")
	print(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.DatabaseConnectionFailure("psql", err)
	} else {
		logger.DatabaseConnectionSuccess("psql")
	}
	defer db.Close()
	h.Register(e.Group("api"))
	port := 1323
	serverError := e.Start(":1323")
	if serverError != nil {
		logger.ServerStartFailure(port, serverError)
	}
}