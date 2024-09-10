package database

import "fmt"

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	row, err := d.DataBaseConnection.Query(query)
	fmt.Println(row.Columns())
	return
}
