package handlers

import (
	"github.com/dowling-john/DjanGopher/http"
	http2 "net/http"
)

type InternalServerError struct {
	BaseHandler
}

func (n *InternalServerError) Connect(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}

func (n *InternalServerError) Delete(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}

func (n *InternalServerError) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusTeapot,
	}
}

func (n *InternalServerError) Head(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}

func (n *InternalServerError) Patch(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}

func (n *InternalServerError) Post(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}

func (n *InternalServerError) Put(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusInternalServerError,
	}
}
