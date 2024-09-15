package http

import (
	"github.com/dowling-john/DjanGopher/config"
	"github.com/dowling-john/DjanGopher/database"
	"log/slog"
	"net/http"
)

// Request
// The request object is built from the raw http incoming request and decorated with useful items for the handlers
// to use
type Request struct {
	Configuration      *config.Configuration
	DataBaseConnection *database.Database
	Logger             *slog.Logger
	BaseHttpRequest    *http.Request
}
