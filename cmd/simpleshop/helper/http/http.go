package http

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

//StringifyHTTPHeader :
func StringifyHTTPHeader(c echo.Context) string {
	m := c.Request().Header
	b := new(bytes.Buffer)
	// Loop over header names
	fmt.Fprintf(b, "[")
	n := len(m)
	i := 0
	for name, values := range m {
		// Loop over all values for the name.
		nn := len(values)
		ii := 0
		for _, value := range values {
			fmt.Fprintf(b, "{\"%s\":\"%s\"}", name, value)
			if ii < nn-1 {
				fmt.Fprintf(b, ",")
			}
			ii++
		}
		if i < n-1 {
			fmt.Fprintf(b, ",")
		}
		i++
	}
	fmt.Fprintf(b, "]")
	return b.String()
}

//FindHTTPHeader :
func FindHTTPHeader(headerName string, c echo.Context) (string, string) {
	var key string
	var val string
	m := c.Request().Header
	b := new(bytes.Buffer)

	// Loop over header names
	for name, values := range m {
		// Loop over all values for the name.
		fmt.Fprintf(b, "{")
		for _, value := range values {
			if name == headerName {
				key = name
				val = value
				break
			}
		}
	}
	return key, val
}

//StringifyHTTPBody :
func StringifyHTTPBody(c echo.Context) string {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}
