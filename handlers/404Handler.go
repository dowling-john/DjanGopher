package handlers

import (
	"github.com/dowling-john/DjanGopher/http"
	http2 "net/http"
)

type NotFoundHandler struct {
	BaseHandler
}

func (n *NotFoundHandler) Connect(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Delete(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Head(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Patch(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Post(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}

func (n *NotFoundHandler) Put(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
	}
}
