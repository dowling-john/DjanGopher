package routing

import "net/http"

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	err := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request)).Write(w)
	if err != nil {
		err := router.InternalServerErrorHandler.Get(request).Write(w)
		if err != nil {
			router.Logger.Fatalf("Error writing response: %v", err)
		}
	}
}
