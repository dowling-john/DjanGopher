package routing

import (
	"github.com/dowling-john/DjanGopher/database"
	"github.com/dowling-john/DjanGopher/handlers"
	"log"
)

type (
	// Router
	// The Router
	Router struct {
		DatabaseConnection         *database.Database
		Logger                     *log.Logger
		Routes                     []*Route
		InternalServerErrorHandler handlers.Handler
		NotFoundHandler            handlers.Handler
	}

	// Route
	// The Route holds the http path and a handler that meets the handlers.Handler interface.
	// The Path variable should be a Regex string that is used to match the path of the incoming request, the selection will
	// return the first match in the routes list, and will return the NotFoundHandler in cases where no route is found.
	Route struct {
		Path    string
		Handler handlers.Handler
	}
)
