package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lmnzr/simpleshop/handler"
)

func main() {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// router.Use(middleware.CORS())

	// CORS restricted
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"https://www.google.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	router.GET("/",handler.Hello)

	router.Start(":9000")
}
