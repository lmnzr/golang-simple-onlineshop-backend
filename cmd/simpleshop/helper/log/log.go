package log

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/config"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/http"
	log "github.com/sirupsen/logrus"
)

func setup() map[string]string{
	configmap := make(map[string]string)

	config, conferr := config.GetConfig()

	if conferr != nil {
		configmap["timeformat"] = "2006-01-02T15:04:05"
		configmap["timezone"] = "Asia/Jakarta"
	} else {
		configmap["timeformat"] = config.GetString("loggerTimeFormat")
		configmap["timezone"] = config.GetString("loggerTimeZone")
	}
	return configmap
	
}

//Logger : Middleware
func Logger(c echo.Context) *log.Entry {
	configmap := setup()
	timeformat := configmap["timeformat"]
	timezone := configmap["timezone"]

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().In(loc).Format(timeformat),
			"tag" : "general",
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
		"tag"	: "http",
	})
}

//LoggerDB : Middleware
func LoggerDB() *log.Entry {
	configmap := setup()
	timeformat := configmap["timeformat"]
	timezone := configmap["timezone"]

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	return log.WithFields(log.Fields{
		"at": time.Now().In(loc).Format(timeformat),
		"tag" : "databse",
	})
}
