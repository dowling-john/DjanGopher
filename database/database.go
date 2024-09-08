package database

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/errors"
)

func InitDatabase(databaseConfiguration config.DatabaseConfig) *sql.DB {
	databaseConnection, err := sql.Open("", databaseConfiguration.FormatDsn())
	errors.LogAnyErrorAndExit(err)
	return databaseConnection
}
