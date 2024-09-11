package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	row := d.DataBaseConnection.QueryRow(query).Scan(&model)
	fmt.Println(row)
	if err != nil {
		return err
	}
	if row == nil {
		return errors.New("no rows found")
	}
	return
}

func ParseSqlRow(row *sql.Rows) (err error) {

	return
}
