package routing

import (
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/dowling-john/DjanGopher/http"
	"github.com/stretchr/testify/assert"
	"log/slog"
	http2 "net/http"
	"net/url"
	"os"
	"testing"
)

type TestHandler struct {
	handlers.BaseHandler
}

// TestSelectHandlerReturnsFirstMatchingRoute
// Ensures that the selectHandler method of the router returns the first regex match
// in the list of routes
func TestSelectHandlerReturnsFirstMatchingRoute(t *testing.T) {
	router := &Router{
		Routes: []*Route{
			{
				Path:    "/test_path",
				Handler: &TestHandler{},
			},
		},
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	handler := router.selectHandler(
		&http.Request{
			BaseHttpRequest: &http2.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/test_path",
				},
			},
			Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		},
	)

	assert.NotNil(t, handler)
	assert.IsType(t, &TestHandler{}, handler)
}

// TestRouterUsesNotFoundHandlerWhenNoRoutesInRouter
// Ensure that the router uses the handler in the NotFoundHandlerVariable if there are no
// routes in the Routes variable.This assumes that the not found handler has been configured.
func TestRouterUsesNotFoundHandlerWhenNoRoutesInRouter(t *testing.T) {

	router := &Router{
		Routes: []*Route{
			{
				Path:    "/test_path",
				Handler: &TestHandler{},
			},
		},
		Logger:          slog.New(slog.NewTextHandler(os.Stdout, nil)),
		NotFoundHandler: &handlers.NotFoundHandler{},
	}

	handler := router.selectHandler(
		&http.Request{
			BaseHttpRequest: &http2.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/not_found_path",
				},
			},
			Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		},
	)

	assert.NotNil(t, handler)
	assert.IsType(t, &handlers.NotFoundHandler{}, handler)
}

// TestRouterUsesNotFoundHandlerWhenNoMatchingRoutesInRouter
// Ensure that the router uses the handler in the NotFoundHandlerVariable if there are no
// matching routes in the Routes variable.This assumes that the not found handler has been configured.
func TestRouterUsesNotFoundHandlerWhenNoMatchingRoutesInRouter(t *testing.T) {

	router := &Router{
		Routes: []*Route{
			{
				Path:    "/test_path",
				Handler: &TestHandler{},
			},
		},
		Logger:          slog.New(slog.NewTextHandler(os.Stdout, nil)),
		NotFoundHandler: &handlers.NotFoundHandler{},
	}

	handler := router.selectHandler(
		&http.Request{
			BaseHttpRequest: &http2.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/not_found_path",
				},
			},
			Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		},
	)

	assert.NotNil(t, handler)
	assert.IsType(t, &handlers.NotFoundHandler{}, handler)
}
