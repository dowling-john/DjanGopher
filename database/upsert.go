package database

import "fmt"

const (
	UpsertQuery = `INSERT INTO %v (%v) VALUES (%v) ON CONFLICT DO UPDATE;`
)

// UpsertOne
// Insert a single model into the database
// This method will update if there is already a record in the DB
func (d *Database) UpsertOne(model interface{}) (err error) {
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
