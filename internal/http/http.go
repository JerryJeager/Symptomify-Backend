package http


type ErrorJson struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func GetErrorJson(err error, message string) *ErrorJson {
	return &ErrorJson{
		Message: message,
		Error:   err.Error(),
	}
}