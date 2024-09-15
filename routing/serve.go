package routing

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	httpResponse := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request))

	if httpResponse.Body == nil {
		httpResponse.Body = io.NopCloser(strings.NewReader(""))
	}

	if _, err := io.Copy(w, httpResponse.Body); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
