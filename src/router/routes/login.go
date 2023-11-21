package routes

import (
	"api.devbook/src/controller"
)

var loginRoute = Route{
	URI:          "/login",
	Method:       "POST",
	Func:         controller.Login,
	RequiresAuth: false,
}
