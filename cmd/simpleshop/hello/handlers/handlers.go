package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello/controllers"
	"net/http"
)

// Hello Endpoint godoc
// @Summary Our Favorite words
// @Produce json
// @Success 200 {object} models.Hello
// @Router /hello [get]
func Hello(context echo.Context) error {
	data := controllers.GetMyHello("Hello World !!!")
	jsondata, _ := json.Marshal(data)
	fmt.Println(string(jsondata))
	return context.JSON(http.StatusOK, data)
}
