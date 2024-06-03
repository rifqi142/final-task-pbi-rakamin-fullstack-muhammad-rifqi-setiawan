package helper

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func CreateResponse(isSuccess bool, data any, errorMessage string, message string) Response {
	return Response{
		Success: isSuccess,
		Data:    data,
		Error:   errorMessage,
		Message: message,
	}
}