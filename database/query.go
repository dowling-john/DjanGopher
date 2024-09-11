package database

import (
	"database/sql"
	"reflect"
)

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	row, err := d.DataBaseConnection.Query(query)
	err = d.buildModel(row, model)
	return
}

func (d *Database) buildModel(row *sql.Rows, model interface{}) (err error) {

	// Get list of column names
	columns, err := row.Columns()

	fieldAddresses := make([]uintptr, len(columns))

	// loop over the fields of the model
	t := reflect.TypeOf(&model)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		modelValue := field.Tag.Get("model")
		for _, column := range columns {
			if column == modelValue {
				fieldAddresses = append(fieldAddresses, field.Offset)
			}
		}
	}
	// scan model tags into struct
	err = row.Scan(fieldAddresses)

	return
}
