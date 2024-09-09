package routing

import "github.com/dowling-john/DjanGopher/handlers"

var InternalServerErrorRoute Route = Route{
	"", &handlers.InternalServerError{},
}
