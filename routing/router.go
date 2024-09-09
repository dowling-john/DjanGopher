package routing

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/handlers"
	"net/http"
)

type (
	// Router
	// The Router
	Router struct {
		DatabaseConnection *sql.DB
		Routes             []*Route
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
	selectedRoute := router.selectRoute()

	err := router.runHttpMethodOfSelectedRoute(router.buildRequest(r), selectedRoute).Write(w)
	if err != nil {
		return
	}
}
