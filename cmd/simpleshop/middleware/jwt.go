package middleware


import (
	// "fmt"
	"strings"
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/http"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/jwt"
)

//JwtMiddleware :
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, val := http.FindHTTPHeader("Authorization", c)
		val = strings.SplitN(val, " ", 2)[1]
		val = strings.TrimSpace(val)

		credential, err := jwt.Parse(val)
		if err != nil {
			return  echo.NewHTTPError(401, "Unathorized")
		}

		c.Set("name", credential.Name)
		c.Set("uuid", credential.UUID)
		c.Set("isAdmin", credential.Admin)
		return next(c)
	}
}