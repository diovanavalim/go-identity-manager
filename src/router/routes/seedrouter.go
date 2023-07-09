package routes

import (
	"identity-manager/src/controller"
	"net/http"
)

var SeedRouter = Route{
	URI:    "/seed",
	Method: http.MethodGet,
	Func:   controller.CreateSeed,
}
