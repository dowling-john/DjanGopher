package routing

import (
	"github.com/dowling-john/DjanGopher/database"
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/stretchr/testify/assert"
	"log/slog"
	http2 "net/http"
	"os"
	"testing"
)

// TestBuildAttachesCorrectElementsToIncomingRequest
// Ensure that the router attaches the correct elements to the incoming requests
func TestBuildAttachesCorrectElementsToIncomingRequest(t *testing.T) {

	router := &Router{
		DatabaseConnection: &database.Database{},
		Routes: []*Route{
			{
				Path:    "/test_path",
				Handler: &TestHandler{},
			},
		},
		Logger:          slog.New(slog.NewTextHandler(os.Stdout, nil)),
		NotFoundHandler: &handlers.NotFoundHandler{},
	}

	wrappedHttpRequest := router.buildRequest(&http2.Request{})

	assert.NotNil(t, wrappedHttpRequest.Logger)
	assert.NotNil(t, wrappedHttpRequest.DataBaseConnection)
}
