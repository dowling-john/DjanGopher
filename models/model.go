package models

import (
	"database/sql"
	"fmt"
	"reflect"
)

type (
	ModelType interface {
		Save() (err error)
		RawQuery(queryString string) (models []ModelType, err error)
		ParseSqlRow(row *sql.Rows) (err error)
	}

	// Model
	// The Model struct should be used as a composition element in the all defined models, this will give the ability
	// to save a changed model to the database, query models from the database a
	// ToDo: Features to Include:
	//  	 Save Model to a relational DataBase
	//       Query Models from a database
	//       Look at how we might look at relationships from table to table
	//       Get the open connection to the database from a centralized service in the application
	Model struct {
		Database  *sql.DB
		TableName string
		Row       *sql.Row
	}
)

func (m *Model) ParseSqlRow(row *sql.Rows) (err error) {
	t := reflect.TypeOf(*m)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonValue := field.Tag.Get("model")
		fmt.Println(jsonValue)
	}
	return
}

// Save Method
// ToDo: The save method needs to use a generalized config file similar to the django effort
func (m *Model) Save() (err error) {
	return
}

// RawQuery
func (m *Model) RawQuery(queryString string) (models []ModelType, err error) {
	return
}
