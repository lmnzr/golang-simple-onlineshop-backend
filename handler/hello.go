package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"github.com/lmnzr/simpleshop/dataobject"
)

// Hello Handler
func Hello(context echo.Context) error {
	data := dataobject.HTTPResponseOk("Hello World !!!")
	jsondata, _ := json.Marshal(data)
	fmt.Println(string(jsondata))
	return context.JSON(http.StatusOK, data)
}
