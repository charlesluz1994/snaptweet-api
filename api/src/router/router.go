package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

// Generate return a router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigureRoutes(r)
}
