package routing

import (
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/dowling-john/DjanGopher/http"
	http2 "net/http"
	"regexp"
)

func (router *Router) selectHandler(request *http.Request) (selectedHandler handlers.Handler) {
	for _, route := range router.Routes {
		if regexp.MustCompile(route.Path).MatchString(request.BaseHttpRequest.URL.Path) {
			return route.Handler
		}
	}
	return router.NotFoundHandler
}

func (router *Router) runHttpMethodOfSelectedHandler(request *http.Request, selectedHandler handlers.Handler) (response *http2.Response) {
	switch request.BaseHttpRequest.Method {
	case http2.MethodGet:
		return selectedHandler.Get(request)
	case http2.MethodPost:
		return selectedHandler.Post(request)
	case http2.MethodPut:
		return selectedHandler.Put(request)
	case http2.MethodPatch:
		return selectedHandler.Patch(request)
	case http2.MethodDelete:
		return selectedHandler.Delete(request)
	default:
		return selectedHandler.Get(request)
	}
}
