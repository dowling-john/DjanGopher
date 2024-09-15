package routing

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	IoCopyError = "error copying io to http writer: %v"
	NilHttpBody = "http body is nil"
)

// ServeHTTP
// Main routing function, this function handles all the incoming http requests and distributes them to the relevant
// handlers.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := router.buildRequest(r)
	httpResponse := router.runHttpMethodOfSelectedHandler(request, router.selectHandler(request))

	if httpResponse.Body == nil {
		router.Logger.Error(NilHttpBody)
		httpResponse.Body = io.NopCloser(strings.NewReader(""))
	}

	if _, err := io.Copy(w, httpResponse.Body); err != nil {
		router.Logger.Error(fmt.Sprintf(IoCopyError, err))
	}
}
