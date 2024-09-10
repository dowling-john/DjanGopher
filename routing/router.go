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
	// The Route holds the http path and a handler that meets the handlers.Handler interface.
	// The Path variable should be a Regex string that is used to match the path of the incoming request, the selection will
	// return the first match in the routes list, and will return the NotFoundHandler in cases where no route is found.
	Route struct {
		Path    string
		Handler handlers.Handler
	}
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	err := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request)).Write(w)
	if err != nil {
		err := router.InternalServerErrorHandler.Get(request).Write(w)
		if err != nil {
			router.Logger.Fatalf("Error writing response: %v", err)
		}
	}
}
