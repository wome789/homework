package main

type Result struct {
	Code    int
	Message string
	Data    any
}

func Ok() Result {
	return Result{Code: 200, Message: ""}
}

func OkWithData(data any) Result {
	return Result{Code: 200, Message: "", Data: data}
}

func failWitMessage(message string) Result {
	return Result{Code: 500, Message: message}
}

func failWithCodeAndMessage(code int, message string) Result {
	return Result{Code: code, Message: message}
}

func failWithCodeAndMessageAndData(code int, message string, data any) Result {
	return Result{Code: code, Message: message, Data: data}
}
