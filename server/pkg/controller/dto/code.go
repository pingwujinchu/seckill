package dto

import "fmt"

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	ParamError          = &Errno{Code: 401, Message: "参数有误"}
	InternalServerError = &Errno{Code: 500, Message: "Internal server error"}
)

//Errno ...
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

//Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

//Error error
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr ...
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
