package database

import (
	"database/sql"
	"djangopher/config"
	"djangopher/errors"
)

func InitDatabase(databaseConfiguration config.DatabaseConfig) *sql.DB {
	databaseConnection, err := sql.Open("", databaseConfiguration.FormatDsn())
	errors.LogAnyErrorAndExit(err)
	return databaseConnection
}
