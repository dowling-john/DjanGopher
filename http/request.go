package http

import (
	"database/sql"
	"github.com/dowling-john/DjanGopher/config"
	"log"
	"net/http"
)

// Request
// The request object is built from the raw http incoming request and decorated with useful items for the handlers
// to use
type Request struct {
	Configuration      *config.Configuration
	DataBaseConnection *sql.DB
	Logger             *log.Logger
	BaseHttpRequest    *http.Request
}
