package database

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	t := reflect.TypeOf(model)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonValue := field.Tag.Get("model")
		fmt.Println(jsonValue)
	}

	row, err := d.DataBaseConnection.Query(query)
	fmt.Println(row.Columns())
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
