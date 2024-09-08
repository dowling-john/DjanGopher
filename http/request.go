package http

import (
	"database/sql"
	"djangopher/config"
	"net/http"
)

// Request
// The request object is built from the raw http incoming request and decorated with useful items for the handlers
// to use
type Request struct {
	Configuration      *config.Configuration
	DataBaseConnection *sql.DB
	BaseHttpRequest    *http.Request
}
