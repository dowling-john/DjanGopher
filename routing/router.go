package routing

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/handlers"
	"net/http"
)

type (
	Router struct {
		DatabaseConnection *sql.DB
		Routes             []*Route
	}

	Route struct {
		Path    string
		Handler handlers.Handler
	}
)

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	selectedRoute, err := router.selectRoute()
	if err != nil {
		return
	}

	err = router.runHttpMethodOfSelectedRoute(router.buildRequest(r), selectedRoute).Write(w)
	if err != nil {
		return
	}
}
