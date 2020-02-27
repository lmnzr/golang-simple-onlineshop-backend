package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lmnzr/simpleshop/cmd/simpleshop/docs"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/env"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		middleware.Logger(nil).Error("Error loading .env file")
	}
	
	environment := env.Getenv("ENVIRONMENT", "development")

	if environment != "development" {
		log.SetLevel(log.InfoLevel)

		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(&lumberjack.Logger{
			Filename:   "var/log/app.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

// @title Simpleshop Swagger API
// @version 1.0
// @description Swagger API for Golang Project Simpleshop.
// @termsOfService http://swagger.io/terms/

// @BasePath
func main() {
	router := echo.New()
	middleware.Setup(router)

	port := env.GetenvI("PORT", 9000)

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	hello.Routes(router)

	lock := make(chan error)
	go func(lock chan error) { lock <- router.Start(fmt.Sprintf(":%d", port)) }(lock)

	err := <-lock
	if err != nil {
		middleware.Logger(nil).Panic("failed to start application")
	}
}
