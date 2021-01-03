package ecode

import (
	"github.com/pkg/errors"
)

var OkCode = 0
var ErrCode = 1

type Code struct {
	Code int
	Message string
}

func New(code int, message string) *Code {
	return &Code{
		Code:    code,
		Message: message,
	}
}

func Error(err error) *Code {
	return &Code{
		Code:    ErrCode,
		Message: err.Error(),
	}
}

func (e *Code) Error() string {
	if e.Code == OkCode {
		return "ok"
	}
	return e.Message
}

func (e *Code) String() string {
	if e.Code == OkCode {
		return "ok"
	}
	return e.Message
}

func Cause(e error) *Code {
	if e == nil {
		return nil
	}
	ec, ok := errors.Cause(e).(*Code)
	if ok {
		return ec
	}
	return Error(e)
}