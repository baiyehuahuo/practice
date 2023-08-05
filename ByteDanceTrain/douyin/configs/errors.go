package configs

import "errors"

var (
	LatestTimeParamError = errors.New("latest time is wrong")
	ParamEmptyError      = errors.New("required param is empty")
)
