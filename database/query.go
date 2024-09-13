package database

import "fmt"

const (
	InsertQuery = `INSERT INTO %v (%v) VALUES (%v)`

	QueryError = "RawQueryError: %v"
)

// InsertOne
// Insert a single model into the database
// This method will error if there is already a record in the DB
func (d *Database) InsertOne(model interface{}) (err error) {
	columns, values, err := d.getModelInsertQueryString(model)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(QueryError, err)
	}

	fmt.Printf(InsertQuery, d.getModelName(model), columns, values)

	_, err = d.DataBaseConnection.Exec(
		fmt.Sprintf(InsertQuery, d.getModelName(model), columns, values),
	)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(QueryError, err)
	}
	return
}

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
