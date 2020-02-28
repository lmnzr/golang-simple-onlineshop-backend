package models

//Hello : Json Format of Hello
type Hello struct {
	Status  int    `json:"s" xml:"s" example:"200"`
	Message string `json:"d" xml:"d" example:"Hello World !!!"`
	Origin  string `json:"o" xml:"o" example:"Default"`
}

//SetStatus :
func (hello *Hello) SetStatus(status int) *Hello {
	hello.Status = status
	return hello
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
