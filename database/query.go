package database

import "fmt"

const (
	InsertQuery = `INSERT INTO %v (%v) VALUES (%v)`

	QueryError = "RawQueryError: %v"
)

// RawQuery
// This method uses the string to query the database and attempt to put the result into the model.
// This method has no checking on the incoming query string so the caller is responsible for checking the
// query for correctness and for security.
func (d *Database) RawQuery(query string, model interface{}) (err error) {
	row, err := d.DataBaseConnection.Query(query)
	if err != nil {
		return fmt.Errorf(QueryError, err)
	}

	err = d.buildModel(row, model)
	if err != nil {
		return fmt.Errorf(QueryError, err)
	}

	return
}
