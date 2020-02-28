package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/http"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/jwt"
)

//JwtMiddleware :
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		valid := false

		_, val := http.FindHTTPHeader("Authorization", c)
		if val != "" {
			parsed := strings.SplitN(val, " ", 2)
			if len(parsed) == 2 {
				if strings.TrimSpace(parsed[0]) == "Bearer"{
					val = parsed[1]
					val = strings.TrimSpace(val)

					credential, err := jwt.Parse(val)
					if err == nil {
						c.Set("name", credential.Name)
						c.Set("uuid", credential.UUID)
						c.Set("isAdmin", credential.Admin)
						valid = true
					} 
				}
			} 
		} 

		if !valid {
			return echo.NewHTTPError(401, "Unathorized")
		}
		return next(c)
	}
}
