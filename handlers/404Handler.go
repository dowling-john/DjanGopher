package handlers

import (
	"github.com/dowling-john/DjanGopher/http"
	"io"
	http2 "net/http"
	"strings"
)

type NotFoundHandler struct {
	BaseHandler
}

func (n *NotFoundHandler) Connect(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Delete(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Head(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Patch(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Post(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func (n *NotFoundHandler) Put(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}
