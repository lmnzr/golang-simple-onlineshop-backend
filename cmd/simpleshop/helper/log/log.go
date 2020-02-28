package log

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/http"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/config"
	"time"
)

//Logger : Middleware
func Logger(c echo.Context) *log.Entry {
	config,conferr := config.GetConfig()

	var timeformat string
	var timezone string

	if conferr != nil {
		timeformat = "2006-01-02T15:04:05"
		timezone = "Asia/Jakarta"
	} else {
		timeformat = config.GetString("loggerTimeFormat")
		timezone = config.GetString("loggerTimeZone")
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().In(loc).Format(timeformat),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().In(loc).Format(timeformat),
		"host":   c.Request().Host,
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"body":   http.StringifyHTTPBody(c),
		"header": http.StringifyHTTPHeader(c),
	})
}