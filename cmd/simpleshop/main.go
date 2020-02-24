package main

import (
	// "net/http"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/lmnzr/simpleshop/cmd/simpleshop/docs"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/middleware"
)

// @title Simpleshop Swagger API
// @version 1.0
// @description Swagger API for Golang Project Simpleshop.
// @termsOfService http://swagger.io/terms/

// @BasePath 
func main() {
	router := echo.New()
	middleware.Setup(router)

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	hello.Routes(router)

	router.Start(":9000")
}
