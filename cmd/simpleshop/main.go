package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lmnzr/simpleshop/cmd/simpleshop/docs"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/env"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/jwt"
	logutil "github.com/lmnzr/simpleshop/cmd/simpleshop/helper/log"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
)

type auth struct {
	Token string `json:"token" xml:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlLmNvbSIsImV4cCI6MTU4Mjg4NjI1NiwiaWF0IjoxNTgyODg1MzU2LCJpc3MiOiJleGFtcGxlLmNvbSIsInN1YiI6ImNsaWVudCIsIm5hbWUiOiJBbG1hcyIsInV1aWQiOiIxMTAzNyIsImlzQWRtaW4iOnRydWV9.FUH8Pu6bBj8_5g6hiUygJD1_swb7k8tvTY-l0EfNk90"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logutil.Logger(nil).Error("Error loading .env file")
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

	router.GET("/", public)
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.GET("/forbidden/", forbidden, middleware.JwtMiddleware)
	router.GET("/credential/", credential)

	hello.Routes(router)

	lock := make(chan error)
	go func(lock chan error) { lock <- router.Start(fmt.Sprintf(":%d", port)) }(lock)

	err := <-lock
	if err != nil {
		logutil.Logger(nil).Panic("failed to start application")
	}
}

func public(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to simpleshop")
}

func forbidden(c echo.Context) error {
	return echo.NewHTTPError(404, "Forbidden Land")
}

func credential(c echo.Context) error {
	cred := jwt.Credential{
		Name:  "Almas",
		UUID:  "11037",
		Admin: true,
	}
	pl := jwt.NewPayload(cred)
	token, _ := jwt.Signing(pl)
	a := auth{
		Token: token,
	}
	return c.JSON(http.StatusOK,a)
}
