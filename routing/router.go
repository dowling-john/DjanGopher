package routing

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/handlers"
	"log"
	"net/http"
)

type (
	// Router
	// The Router
	Router struct {
		DatabaseConnection         *sql.DB
		Logger                     *log.Logger
		Routes                     []*Route
		InternalServerErrorHandler handlers.Handler
		NotFoundHandler            handlers.Handler
	}

	// Route
	// The Route holds the http path and a handler that meets the handlers.Handler interface, the routes are added to
	// router object and are used to match the incoming request and direct them to the relevant handler type
	Route struct {
		Path    string
		Handler handlers.Handler
	}
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := router.runHttpMethodOfSelectedHandler(router.buildRequest(r), router.selectHandler()).Write(w)
	if err != nil {

	}
}
