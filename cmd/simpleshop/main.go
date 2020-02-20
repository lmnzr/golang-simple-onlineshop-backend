package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lmnzr/simpleshop/cmd/simpleshop/docs"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/handler"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Simpleshop Swagger API
// @version 1.0
// @description Swagger API for Golang Project Simpleshop.
// @termsOfService http://swagger.io/terms/

// @BasePath /api
func main() {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// router.Use(middleware.CORS())

	// CORS restricted
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.GET("/api/", handler.Hello)

	router.Start(":9000")
}
