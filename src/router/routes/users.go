package routes

import (
	"net/http"

	"api.devbook/src/controller"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controller.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Func:         controller.GetAllUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodGet,
		Func:         controller.GetUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodPut,
		Func:         controller.UpdateUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodDelete,
		Func:         controller.DeleteUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}/follow",
		Method:       http.MethodPost,
		Func:         controller.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}/unfollow",
		Method:       http.MethodDelete,
		Func:         controller.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}/followers",
		Method:       http.MethodGet,
		Func:         controller.GetFollowers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}/following",
		Method:       http.MethodGet,
		Func:         controller.GetFollowing,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{id}/updatePassword",
		Method:       http.MethodPost,
		Func:         controller.UpdatePassword,
		RequiresAuth: true,
	},
}
