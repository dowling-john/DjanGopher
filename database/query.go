package database

import (
	"errors"
)

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	row, err := d.DataBaseConnection.Query(query)
	if err != nil {
		return err
	}
	if row == nil {
		return errors.New("no rows found")
	}
	return
}
