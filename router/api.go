package router

import (
	"example.com/test1/handler"
	"example.com/test1/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	// Route không yêu cầu xác thực JWT
	api.Echo.POST("/user/sign-in", api.UserHandler.HandlerSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
