package routing

import (
	"net/http"
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	httpResponse := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request))

	_ = httpResponse.Write(w)

}
