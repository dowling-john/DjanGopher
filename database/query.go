package database

import (
	"database/sql"
)

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	return d.DataBaseConnection.QueryRow(query).Scan(&model)
}

func ParseSqlRow(row *sql.Rows) (err error) {

	return
}
