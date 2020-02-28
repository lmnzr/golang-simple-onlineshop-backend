package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/log"
)

//LogMiddleware :
func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Logger(c).Info("incoming request")
		return next(c)
	}
}