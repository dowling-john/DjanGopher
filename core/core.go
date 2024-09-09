package core

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/routing"
	"log"
	"net/http"
)

type DjanGopher struct {
	Configuration      *config.Configuration
	DatabaseConnection *sql.DB
	Router             *routing.Router
}

// RunServer
// This is the entry point to the application
//
// The RunServer function serves to initialize the application.
// It will load the configuration, build your routing, and then will handle the incoming http requests.
//
// The Initialization process of the app is as follows.
// - Parse the configuration into an object
// - Initiate a Database connection
// - Add the database connection to the initialized router
func (d *DjanGopher) RunServer() {

	d.Configuration = config.InitConfiguration()
	// ToDo: Add back once the routing is operational
	//d.DatabaseConnection = database.InitDatabase(d.Configuration.DatabaseConfiguration)
	//d.Router.DatabaseConnection = d.DatabaseConnection

	log.Fatal(
		http.ListenAndServe(":8080", d.Router),
	)
}
