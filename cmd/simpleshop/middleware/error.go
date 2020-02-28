package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/log"
)

//StandardError :
type StandardError struct {
	Status  int    `json:"s" xml:"s"`
	Message string `json:"m" xml:"m"`
}

//NewStandardError :
func NewStandardError(message string, code int, router echo.Context) {
	error := &StandardError{
		Status:  code,
		Message: message,
	}
	router.JSON(code, error)
}

//Error : Middleware
func Error(err error, router echo.Context) {
	report, ok := err.(*echo.HTTPError)
	var message string
	var code int
	if ok {
		message = report.Message.(string)
		code = report.Code
	} else {
		message = err.Error()
		code = 400
	}

	logmessage := fmt.Sprintf("%d - %v", code, message)
	log.Logger(router).Error(logmessage)
	NewStandardError(message, code, router)
}
