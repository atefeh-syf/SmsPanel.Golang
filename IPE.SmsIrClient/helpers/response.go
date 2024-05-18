package helpers

import "net/http"

type ResultCode int

const (
	Success         ResultCode = 0
	ValidationError ResultCode = 40001
	AuthError       ResultCode = 40101
	ForbiddenError       ResultCode = 40301
	NotFoundError ResultCode = 40401
	LimiterError    ResultCode = 42901
	OtpLimiterError ResultCode = 42902
	CustomRecovery  ResultCode = 50001
	InternalError   ResultCode = 50002
)

type Response struct {
	Status				int 				`json:"status"`
	Message			string			 `json:"message"`
	Data				 any				`json:"data"`
	//Error				 any				`json:"error"`
}

func GenerateResponse(Status int, message string, data any) *Response {
	return &Response{
		Status: Status,
		Message:    message,
		Data: data,
	}
}


func GenerateResponseWithAnyError(Status int, message string, data any) *Response {
	return &Response{
		Status: Status,
		Message:    message,
		Data: data,
		//Error:      err,
	}
}

var StatusCodeMapping = map[string]int{}

func TranslateErrorToStatusCode(err error) int {
	v, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return v
}