package models

//Hello : Json Format of Hello
type Hello struct {
	Status int 		`json:"s" xml:"s" example:"200"`
	Hello string    `json:"d" xml:"d" example:"Hello World !!!"`
}