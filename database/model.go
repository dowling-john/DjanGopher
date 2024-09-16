package database

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	ModelStructTag = "model"
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

	fieldPointers, err := d.getStructFieldPointers(columns, model_value)
	if err != nil {
		return err
	}

	for row.Next() {
		if err = row.Scan(fieldPointers...); err != nil {
			return err
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
			fieldTag := modelValue.Type().Field(modelFieldIndex).Tag.Get(ModelStructTag)
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

// GetModelName
// Does exactly what it says on the tin !!
func (d *Database) getModelName(model interface{}) (modelName string) {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	fmt.Println(modelValue.Type().String())
	name := strings.Split(".", modelValue.Type().String())
	fmt.Println(name)
	return name[len(name)-1]
}

// getModelInsertQueryString
// This method translates the model into a string seperated by ", " that represents the values and a string seperated by
// ", " that represents the column name, the column name is taken from the struct tag.
//
// ToDo: I am not sure that this function is the best implementation of this I think that we can do better and refactor this
func (d *Database) getModelInsertQueryString(model interface{}) (columnNames, insertQueryString string, err error) {
	var fields []string
	var columns []string
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	for modelFieldIndex := 0; modelFieldIndex < modelValue.NumField(); modelFieldIndex++ {
		if fieldTag := modelValue.Type().Field(modelFieldIndex).Tag.Get(ModelStructTag); fieldTag != "" {
			columns = append(columns, fieldTag)
			// ToDo: not sure that this is the most efficient way to do this.
			switch modelValue.Field(modelFieldIndex).Kind() {
			case reflect.String:
				fields = append(fields, "'"+modelValue.Field(modelFieldIndex).Interface().(string)+"'")
			case reflect.Int:
				fields = append(fields, strconv.Itoa(modelValue.Field(modelFieldIndex).Interface().(int)))
			default:
				err = errors.New(fmt.Sprintf("model contains an unrecognised type of %v", modelValue.Field(modelFieldIndex).String()))
				break
			}
		}
	}
	columnNames = strings.Join(columns, ",")
	insertQueryString = strings.Join(fields, ", ")
	return
}

func (d *Database) getStructFieldPointers(columns []string, modelValue reflect.Value) (pointers []interface{}, err error) {
	for columnIndex := 0; columnIndex < len(columns); columnIndex++ {
		for modelFieldIndex := 0; modelFieldIndex < modelValue.NumField(); modelFieldIndex++ {
			fieldTag := modelValue.Type().Field(modelFieldIndex).Tag.Get(ModelStructTag)
			if fieldTag == columns[columnIndex] {
				pointers = append(pointers, modelValue.Field(modelFieldIndex).Addr().Interface())
			}
		}
	}
	return pointers, nil
}
