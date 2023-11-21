package routes

import (
	"net/http"

	"api.devbook/src/controller"
)

var publicationsRoutes = []Route{
	{
		URI:          "/publications",
		Method:       http.MethodPost,
		Func:         controller.CreatePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications",
		Method:       http.MethodGet,
		Func:         controller.GetAllPublications,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodGet,
		Func:         controller.GetPublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodPut,
		Func:         controller.UpdatePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodDelete,
		Func:         controller.DeletePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/publications",
		Method:       http.MethodGet,
		Func:         controller.GetAllPublicationsOfUser,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}/like",
		Method:       http.MethodPost,
		Func:         controller.LikePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}/dislike",
		Method:       http.MethodPost,
		Func:         controller.DislikePublication,
		RequiresAuth: true,
	},
}
