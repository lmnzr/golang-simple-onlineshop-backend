package main

import (
	"github.com/labstack/echo"
	"simpleshop/handler"
)

func main() {
	router := echo.New()

	router.GET("/",handler.Hello)


	router.Start(":9000")
}
