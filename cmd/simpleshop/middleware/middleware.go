package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Setup : Middleware
func Setup(router *echo.Echo) {
	router.Use(Session())
	router.Use(echo.WrapMiddleware(CORS().Handler))
	router.HTTPErrorHandler = Error
	router.Use(LogMiddleware)
	router.Use(middleware.Recover())
}
