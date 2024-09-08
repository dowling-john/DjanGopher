package handlers

import (
	"djangopher/http"
	http2 "net/http"
)

type (
	// Handler
	// The Handler Interface
	Handler interface {
		Connect(request *http.Request) (httpResponse *http2.Response)
		Delete(request *http.Request) (httpResponse *http2.Response)
		Get(request *http.Request) (httpResponse *http2.Response)
		Head(request *http.Request) (httpResponse *http2.Response)
		Patch(request *http.Request) (httpResponse *http2.Response)
		Post(request *http.Request) (httpResponse *http2.Response)
		Put(request *http.Request) (httpResponse *http2.Response)
	}

	BaseHandler struct {
	}
)

// Connect
// Basic implementation of the Connect Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Connect(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Delete
// Basic implementation of the Delete Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Delete(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Get
// Basic implementation of the Get Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Get(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Head
// Basic implementation of the Head Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Head(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Patch
// Basic implementation of the Patch Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Patch(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Post
// Basic implementation of the Post Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Post(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}

// Put
// Basic implementation of the Put Method, This method returns a Status of 405 (MethodNotAllowed)
// You can choose not to implement any of the responses we will take care or the appropriate response
func (b *BaseHandler) Put(request *http.Request) (httpResponse *http2.Response) {
	return &http2.Response{
		StatusCode: http2.StatusMethodNotAllowed,
	}
}
