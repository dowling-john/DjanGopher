package database

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/errors"
)

// Database
// Todo: need to add the logger to the database
type Database struct {
	DataBaseConnection *sql.DB
}

func InitDatabase(databaseConfiguration config.DatabaseConfig) *Database {
	databaseConnection, err := sql.Open(databaseConfiguration.DriverName, databaseConfiguration.DatabaseName)
	errors.LogAnyErrorAndExit(err)
	return &Database{
		DataBaseConnection: databaseConnection,
	}
}
