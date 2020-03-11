package models

//Hello : Json Format of Hello
type Hello struct {
	Message string `json:"message" example:"Hello World !!!"`
	Origin  string `json:"origin" example:"Default"`
}

//SetMessage :
func (hello *Hello) SetMessage(message string) *Hello {
	hello.Message = message
	return hello
}

//SetOrigin :
func (hello *Hello) SetOrigin(origin string) *Hello {
	hello.Origin = origin
	return hello
}
