package routing

import (
	"djangopher/http"
	http2 "net/http"
)

func (router *Router) selectRoute() (selectedRoute *Route, err error) {
	return
}

func (router *Router) runHttpMethodOfSelectedRoute(request *http.Request, selectedRoute *Route) (response *http2.Response) {
	switch request.BaseHttpRequest.Method {
	case http2.MethodGet:
		return selectedRoute.Handler.Get(request)
	case http2.MethodPost:
		return selectedRoute.Handler.Post(request)
	case http2.MethodPut:
		return selectedRoute.Handler.Put(request)
	case http2.MethodPatch:
		return selectedRoute.Handler.Patch(request)
	case http2.MethodDelete:
		return selectedRoute.Handler.Delete(request)
	default:
		return selectedRoute.Handler.Get(request)
	}
}
