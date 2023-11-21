package routes

import (
	"net/http"

	"api.devbook/src/middleware"
	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Config coloca sobe todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, publicationsRoutes...)

	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Auth(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
