package router

import (
	"api.devbook/src/router/routes"
	"github.com/gorilla/mux"
)

// Gerar retorna um router com as rotas configuradas
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
