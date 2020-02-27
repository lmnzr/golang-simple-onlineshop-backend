package middleware

import (
	b64 "encoding/base64"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//Session : Middleware
func Session() echo.MiddlewareFunc {
	
	config,conferr := config.GetConfig()

	var sessionsecret string

	if conferr != nil {
		sessionsecret = "secret"
	} else {
		sessionsecret = config.GetString("sessionSecret")
	}

	sEnc := b64.StdEncoding.EncodeToString([]byte(sessionsecret))
	return session.Middleware(sessions.NewCookieStore([]byte(sEnc)))
}
