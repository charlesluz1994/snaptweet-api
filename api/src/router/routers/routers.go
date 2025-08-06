package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route Router represent all apis routes.
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// ConfigureRoutes sets all routes inside router
// Receive a router with none route, and populate it.
func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := userRouters

	//insert routes inside router.
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		routes = append(routes, route)
	}

	return r
}
