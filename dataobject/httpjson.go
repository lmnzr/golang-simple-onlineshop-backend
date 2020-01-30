package dataobject

// HTTPResponse : Json Format of HTTP Response
type HTTPResponse struct {
	Status int    `json:"s" xml:"s"`
	Data   string `json:"d" xml:"d"`
}

// HTTPResponseError : Json Format of HTTP Response With Error Debug
type HTTPResponseError struct {
	Status int    `json:"s" xml:"s"`
	Data   string `json:"d" xml:"d"`
	Error  string `json:"e" xml:"e"`
}

//HTTPResponseOk : Function Return JSON Formatted Ok Response
func HTTPResponseOk(data string) *HTTPResponse {
	json := &HTTPResponse{
		Status: 200,
		Data:   data,
	}
	return json
}

//HTTPResponseFail : Function Return JSON Formatted Fail Response
func HTTPResponseFail(data string) *HTTPResponse {
	json := &HTTPResponse{
		Status: 500,
		Data:   data,
	}
	return json
}

//HTTPResponseFailError : Function Return JSON Formatted Fail Response With Error Debug
func HTTPResponseFailError(data string, debug string) *HTTPResponseError {
	json := &HTTPResponseError{
		Status: 500,
		Data:   data,
		Error:  debug,
	}
	return json
}

//HTTPResponseNotFound : Function Return JSON Formatted Not Found Response
func HTTPResponseNotFound() *HTTPResponse {
	json := &HTTPResponse{
		Status: 404,
		Data:   "Resource Not Found",
	}
	return json
}
