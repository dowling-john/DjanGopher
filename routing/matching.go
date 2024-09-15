package routing

import (
	"fmt"
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/dowling-john/DjanGopher/http"
	http2 "net/http"
	"regexp"
)

const (
	IncomingRequestLogging      = "incoming request for path: %v"
	MatchedRouteLogging         = "matching route: %v"
	MatchedRouteNotFoundLogging = "matching route not found: %v"
	MatchedHttpMethod           = "matching http method: %v"
	UnknownHttpMethod           = "unknown http method: %v"
)

func (router *Router) selectHandler(request *http.Request) (selectedHandler handlers.Handler) {
	router.Logger.Debug(fmt.Sprintf(IncomingRequestLogging, request.BaseHttpRequest.URL.Path))
	for _, route := range router.Routes {
		if regexp.MustCompile(route.Path).MatchString(request.BaseHttpRequest.URL.Path) {
			router.Logger.Debug(fmt.Sprintf(MatchedRouteLogging, route.Path))
			return route.Handler
		}
	}
	router.Logger.Debug(fmt.Sprintf(MatchedRouteNotFoundLogging, request.BaseHttpRequest.URL.Path))
	return router.NotFoundHandler
}

// runHttpMethodOfSelectedHandler
// Run the correct method of the selected route handler and return the resulting http response to the serve method
func (router *Router) runHttpMethodOfSelectedHandler(request *http.Request, selectedHandler handlers.Handler) (response *http2.Response) {
	if selectedHandler != nil {
		switch request.BaseHttpRequest.Method {
		case http2.MethodGet:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Get(request)
		case http2.MethodPost:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Post(request)
		case http2.MethodPut:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Put(request)
		case http2.MethodPatch:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Patch(request)
		case http2.MethodDelete:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Delete(request)
		default:
			router.Logger.Debug(fmt.Sprintf(MatchedHttpMethod, request.BaseHttpRequest.Method))
			return selectedHandler.Get(request)
		}
	}
	router.Logger.Error(fmt.Sprintf(UnknownHttpMethod, request.BaseHttpRequest.Method))
	errorHandler := &handlers.InternalServerError{}
	return errorHandler.Get(request)
}
