package hello

import (
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello/handlers"
)

// Routes : All Hello Routing
func Routes(router *echo.Echo) {
	// Routes
	hello := router.Group("/hello")
	{
		hello.GET("/", handlers.GetHello)
		hello.POST("/", handlers.PostHello)
	}
}
