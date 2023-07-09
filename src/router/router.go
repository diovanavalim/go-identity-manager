package router

import (
	"github.com/gorilla/mux"
	"identity-manager/src/router/routes"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return ConfigRouter(router)
}

func ConfigRouter(r *mux.Router) *mux.Router {
	var pkgRoutes []routes.Route

	pkgRoutes = append(pkgRoutes, routes.DIDRoute)
	pkgRoutes = append(pkgRoutes, routes.SeedRouter)

	for _, route := range pkgRoutes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
