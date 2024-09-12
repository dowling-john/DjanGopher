package database

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

const (
	MODELSTRUCTTAG = "model"
)

// buildModel
// This takes the incoming model and the match Row and maps the columns of the sql query to the
// fields of the model struct to the columns in the row.
func (d *Database) buildModel(row *sql.Rows, model interface{}) (err error) {
	columns, err := row.Columns()

	model_value := reflect.ValueOf(model)
	if model_value.Kind() == reflect.Ptr {
		model_value = model_value.Elem()
	}

	if err = d.checkStructContainsCorrectFieldsToContainRow(columns, model_value); err != nil {
		return err
	}

	fieldPointers, err := d.GetStructFieldPointers(columns, model_value)
	if err != nil {
		return err
	}

	for row.Next() {
		err = row.Scan(fieldPointers...)
		if err != nil {
			break
		}
	}

	return
}

// checkStructContainsCorrectFieldsToContainRow
// ensure that the struct contains the correct number and the correct tag names to get the column values from the sql
// row
func (d *Database) checkStructContainsCorrectFieldsToContainRow(columns []string, modelValue reflect.Value) (err error) {
	for columnIndex := 0; columnIndex < len(columns); columnIndex++ {
		columnFound := false
		for modelFieldIndex := 0; modelFieldIndex < modelValue.NumField(); modelFieldIndex++ {
			fieldTag := modelValue.Type().Field(modelFieldIndex).Tag.Get(MODELSTRUCTTAG)
			fmt.Println(fieldTag, columns[columnIndex])
			if fieldTag == columns[columnIndex] {
				columnFound = true
				break
			}
		}
		if !columnFound {
			return errors.New(fmt.Sprintf("model does not contain field %v", columns[columnIndex]))
		}
	}
	return nil
}

func (d *Database) GetStructFieldPointers(columns []string, modelValue reflect.Value) (pointers []interface{}, err error) {
	for columnIndex := 0; columnIndex < len(columns); columnIndex++ {
		for modelFieldIndex := 0; modelFieldIndex < modelValue.NumField(); modelFieldIndex++ {
			fieldTag := modelValue.Type().Field(modelFieldIndex).Tag.Get(MODELSTRUCTTAG)
			if fieldTag == columns[columnIndex] {
				pointers = append(pointers, modelValue.Field(modelFieldIndex).Addr().Interface())
			}
		}
	}
	return pointers, nil
}
