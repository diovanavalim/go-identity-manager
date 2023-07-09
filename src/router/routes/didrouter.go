package routes

import (
	"identity-manager/src/controller"
	"net/http"
)

var DIDRoute = Route{
	URI:    "/did",
	Method: http.MethodGet,
	Func:   controller.CreateDID,
}
