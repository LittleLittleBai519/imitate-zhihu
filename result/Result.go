package result

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

// Usage:
// Use WithData() to bring normal errors,
// In repository layer, DO NOT bring models,
// bring DTOs in service layer instead.
// Use HandleError() to print stack when bringing error as Data.

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res Result) Error() string {
	return strconv.Itoa(res.Code) + ":" + res.Message
}


func (res Result) IsOK() bool {
	return res.Code == 0
}

func (res Result) NotOK() bool {
	return res.Code != 0
}

func (res Result) WithData(data interface{}) Result {
	res.Data = data
	return res
}

func (res Result) WithDataIfOK(data interface{}) Result {
	if res.NotOK() {
		return res
	}
	return res.WithData(data)
}

func (res Result) WithDataError(err error) Result {
	res.Data = err.Error()
	return res
}

func HandleServerErr(err error) Result {
	ers := errors.WithStack(err)
	fmt.Printf("%+v\n", ers)
	return ServerErr.WithDataError(err)
}

// Result definitions
var (
	Ok         = Result{Code: 0, Message: "OK"}
	BadRequest = Result{Code: 400, Message: "Bad Request"}
	ServerErr  = Result{Code: 500, Message: "Server Error"}
	// 100x: Authorization
	EmptyAuth     = Result{Code: 1001, Message: "Empty Authorization"}
	AuthFormatErr = Result{Code: 1002, Message: "Authorization Format Error"}
	TokenErr      = Result{Code: 1003, Message: "Token Error"}
	// 200x: User Login
	UserNotFoundErr       = Result{Code: 2001, Message: "User Not Found"}
	PasswordNotCorrectErr = Result{Code: 2002, Message: "Password not correct"}
	// 201x: User Register
	EmailAlreadyExistErr = Result{Code: 2011, Message: "Email already exists"}
	CreateUserErr        = Result{Code: 2012, Message: "DB Create user error"}
	// 210x: Get Question
	QuestionNotFoundErr = Result{Code: 2101, Message: "Question Not Found"}
	// 211x: Create Question
	CreateQuestionErr = Result{Code: 2111, Message: "DB Create question error"}
)
