package configs

import "errors"

var (
	ParamInputTypeError           = errors.New("input param type is wrong")
	ParamInputLengthExceededError = errors.New("input param length exceeded")
	ParamEmptyError               = errors.New("required param is empty")
	ParamUnknownActionTypeError   = errors.New("unknown action type")
)

var (
	DBCreateUserError = errors.New("create user fail")
)
