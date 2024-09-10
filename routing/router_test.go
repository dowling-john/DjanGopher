package routing

import (
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/dowling-john/DjanGopher/http"
	"reflect"
	"testing"
)

// TestRouterUsesNotFoundHandlerWhenNoRoutesInRouter
// Ensure that the router uses the handler in the NotFoundHandlerVariable if there are no
// routes in the Routes variable.
func TestRouterUsesNotFoundHandlerWhenNoRoutesInRouter(t *testing.T) {

	router := &Router{
		Routes:          []*Route{},
		NotFoundHandler: &handlers.NotFoundHandler{},
	}

	selectedHandler := router.selectHandler(&http.Request{})

	if reflect.TypeOf(selectedHandler).String() != "*handlers.NotFoundHandler" {
		t.Fail()
	}

}
