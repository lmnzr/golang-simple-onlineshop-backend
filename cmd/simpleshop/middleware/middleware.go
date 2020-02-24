package middleware

import (
	"github.com/labstack/echo/v4"
)

//Setup : Middleware
func Setup(router *echo.Echo)  {
	Logger(router)
	CORS(router)
	Error(router)
}