package database

func (d *Database) RawQuery(query string, model interface{}) (err error) {
	d.DataBaseConnection.Query("SELECT * FROM 'testing' ")
	return
}
