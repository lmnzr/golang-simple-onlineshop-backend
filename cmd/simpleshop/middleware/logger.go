package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Logger : Middleware
func Logger(router *echo.Echo)  {
	router.Use(middleware.Logger())
}
