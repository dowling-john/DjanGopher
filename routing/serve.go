package routing

import (
	"io"
	"net/http"
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
// ToDo: This needs a little clean up I think that an error handler might be in order
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	httpResponse := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request))

	w.WriteHeader(httpResponse.StatusCode)

	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		err := router.InternalServerErrorHandler.Get(request).Write(w)
		if err != nil {
			router.Logger.Fatalf("Error writing response: %v", err)
		}
	}

	_, err = w.Write(b)
	if err != nil {
		err := router.InternalServerErrorHandler.Get(request).Write(w)
		if err != nil {
			router.Logger.Fatalf("Error writing response: %v", err)
		}
	}

	if err != nil {
		err := router.InternalServerErrorHandler.Get(request).Write(w)
		if err != nil {
			router.Logger.Fatalf("Error writing response: %v", err)
		}
	}
}
