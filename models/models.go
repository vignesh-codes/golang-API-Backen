package models

import "fmt"

type TodoModel struct {
	Id        int
	Task      string
	Desc      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

type ErrorHandler struct {
	ErrorType    string
	ErrorMessage string
}

type ResponseHandler struct {
	Message string
	Status  int
}

func (e *ErrorHandler) Error() string {
	return fmt.Sprintf("Error Type: %s \n Error Message: ", e.ErrorType, e.ErrorMessage)
}
