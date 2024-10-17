package main

import (
	"github.com/dowling-john/DjanGopher/core"
	"github.com/dowling-john/DjanGopher/handlers"
	"github.com/dowling-john/DjanGopher/http"
	"github.com/dowling-john/DjanGopher/routing"
	_ "github.com/mattn/go-sqlite3"
	"io"
	http2 "net/http"
	"strings"
)

type (
	HelloWorld struct {
		handlers.BaseHandler
	}
)

func (h *HelloWorld) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		Body: io.NopCloser(strings.NewReader("Hello, world!")),
	}
}

func main() {

	app := core.DjanGopher{
		Router: &routing.Router{
			Routes: []*routing.Route{
				{Path: "/", Handler: &HelloWorld{}},
			},
		},
	}

	app.RunServer()
}
