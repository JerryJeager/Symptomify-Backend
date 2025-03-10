package http

import "context"

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

func GetUserID(ctx context.Context) (string, error) {
	return ctx.Value("user_id").(string), nil
}
