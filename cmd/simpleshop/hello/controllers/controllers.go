package controllers

import "github.com/lmnzr/simpleshop/cmd/simpleshop/hello/models"

//GetMyHello : Function Return MyHello
func GetMyHello(helloWord string) *models.Hello {
	json := &models.Hello{
		Status: 200,
		Hello:   helloWord,
	}
	return json
}
