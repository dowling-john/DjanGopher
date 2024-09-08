package handlers

import (
	"djangopher/handlers"
	"djangopher/http"
	http2 "net/http"
)

type Whoo struct {
	handlers.BaseHandler
}

func (w *Whoo) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: 200,
	}
}

func (w *Whoo) Post(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{}
}
