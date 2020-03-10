package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//HTTPResponse : Json Format of Hello Endpoint HTTP Response
type HTTPResponse struct {
	Status int 		`json:"s" xml:"s" example:"200"`
	Data string    	`json:"d" xml:"d" example:"Hello World !!!"`
}

//HTTPResponseOk : Function Return JSON Formatted Hello Endpoint Ok Response
func HTTPResponseOk(data string) *HTTPResponse {
	json := &HTTPResponse{
		Status: 200,
		Data:   data,
	}
	return json
}

// Hello Endpoint godoc
// @Summary Our Favorite words
// @Produce json
// @Success 200 {object} HTTPResponse
// @Router / [get]
func Hello(context echo.Context) error {
	data := HTTPResponseOk("Hello World !!!")
	return context.JSON(http.StatusOK, data)
}
