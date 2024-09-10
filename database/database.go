package database

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/errors"
)

type Database struct {
	DataBaseConnection *sql.DB
}

func InitDatabase(databaseConfiguration config.DatabaseConfig) *Database {
	databaseConnection, err := sql.Open(databaseConfiguration.DriverName, "./foo.db")
	errors.LogAnyErrorAndExit(err)
	return &Database{
		DataBaseConnection: databaseConnection,
	}
}
