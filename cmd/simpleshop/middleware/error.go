package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
	report, _ := err.(*echo.HTTPError)
	message := fmt.Sprintf("%v", report.Message)
	report.Message = fmt.Sprintf("%d - %v", report.Code, report.Message)
	NewStandardError(message, report.Code, router)
	Logger(router).Error(report.Message)
}
