package models

import "database/sql"

type (
	ModelType interface {
		Save() (err error)
		Query() (models []ModelType, err error)
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
	}
)

// Save Method
// ToDo: The save method needs to use a generalized config file similar to the django effort
func (m *Model) Save() (err error) {
	return
}
