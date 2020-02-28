package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

//CORS : Middleware
func CORS() *cors.Cors {
	// CORS restricted
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowedHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
		// Debug: true,
	})
	return corsMiddleware
}
