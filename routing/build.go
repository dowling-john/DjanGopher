package routing

import (
	"djangopher/http"
	http2 "net/http"
)

func (router Router) buildRequest(request *http2.Request) *http.Request {
	return &http.Request{
		DataBaseConnection: router.DatabaseConnection,
		BaseHttpRequest:    request,
	}
}
