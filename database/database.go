package database

import (
	"database/sql"
	"fmt"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/errors"
)

type Database struct {
	DataBaseConnection *sql.DB
}

func InitDatabase(databaseConfiguration config.DatabaseConfig) *Database {
	databaseConnection, err := sql.Open(databaseConfiguration.DriverName, databaseConfiguration.DatabaseName)
	errors.LogAnyErrorAndExit(err)
	fmt.Println(err)
	return &Database{
		DataBaseConnection: databaseConnection,
	}
}
