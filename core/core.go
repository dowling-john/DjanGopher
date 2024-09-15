package core

import (
	"fmt"
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/database"
	"github.com/dowling-john/DjanGopher/logging"
	"github.com/dowling-john/DjanGopher/routing"
	"log/slog"
	"net/http"
)

type DjanGopher struct {
	Configuration      *config.Configuration
	DatabaseConnection *database.Database
	Logger             *slog.Logger
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
	d.Logger = logging.Init(d.Configuration.LoggingConfiguration)
	d.Router.Logger = d.Logger

	d.DatabaseConnection = database.InitDatabase(d.Configuration.DatabaseConfiguration)
	d.Router.DatabaseConnection = d.DatabaseConnection

	if err := http.ListenAndServe(fmt.Sprintf(":%v", d.Configuration.ServerConfiguration.Port), d.Router); err != nil {
		d.Logger.Error("Error starting server", err)
	}
}
