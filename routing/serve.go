package routing

import (
	"io"
	"net/http"
)

// WriteErrorIfRequired
// This method checks for an error and writes this to the response writer, then returns true if it has done something
// and false if not.
func (router *Router) WriteErrorIfRequired(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, err.Error())
		return true
	}
	return false
}

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers
// ToDo: This needs a little clean up I think that an error handler might be in order
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	httpResponse := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request))

	w.WriteHeader(httpResponse.StatusCode)
	b, err := io.ReadAll(httpResponse.Body)
	if !router.WriteErrorIfRequired(w, err) {
		_, err := w.Write(b)
		router.WriteErrorIfRequired(w, err)
	}
}
