package handlers

import (
	// "encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello/controllers"
	"net/http"
)

// GetHello Endpoint godoc
// @Summary Get Our Favorite words
// @Produce json
// @Success 200 {object} models.Hello
// @Router /hello/ [get]
func GetHello(context echo.Context) error {

	fmt.Println(context.Get("name"))
	data := controllers.GetMyHello("Hello World !!!", context)
	// jsondata, _ := json.Marshal(data)
	// fmt.Println(string(jsondata))
	return context.JSON(http.StatusOK, data)
}

// PostHello Endpoint godoc
// @Summary Save Our Favorite words
// @Produce json
// @Success 200 {object} models.Hello
// @Param m body string true "Your Own Hello Word"
// @Param o body string true "Your Own Hello Word Signature"
// @Router /hello/ [post]
func PostHello(context echo.Context) error {
	data := controllers.PostMyHello(context)
	return context.JSON(http.StatusOK, data)
}
