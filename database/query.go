package database

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	_, err = d.DataBaseConnection.Query(query)
	return
}
