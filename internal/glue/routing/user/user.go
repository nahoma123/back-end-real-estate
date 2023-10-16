package user

import (
	"visitor_management/internal/glue/routing"
	"visitor_management/internal/handler/middleware"
	"visitor_management/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup, handler rest.User, authMiddleware middleware.AuthMiddleware) {
	userRoutes := []routing.Router{
		{
			Method:      "POST",
			Path:        "/users/register",
			Handler:     handler.Register,
			Middlewares: []gin.HandlerFunc{},
		},

		// update user
		{
			Method:      "PATCH",
			Path:        "/users/:id",
			Handler:     handler.UpdateUser,
			Middlewares: []gin.HandlerFunc{},
		},
		// get user
		{
			Method:      "GET",
			Path:        "/users/:id",
			Handler:     handler.GetUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "POST",
			Path:        "/auth/login",
			Handler:     handler.Login,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:  "GET",
			Path:    "/users",
			Handler: handler.GetUsers,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
	}
	routing.RegisterRoutes(router, userRoutes)
}
