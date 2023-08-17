package dyerror

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

type DouyinError struct {
	ErrCode    int32
	ErrMessage string
}

var (
	// ParamError 1xx
	ParamInputTypeError           = &DouyinError{ErrCode: 101, ErrMessage: "input param type is wrong"}
	ParamInputLengthExceededError = &DouyinError{ErrCode: 102, ErrMessage: "input param length exceeded"}
	ParamEmptyError               = &DouyinError{ErrCode: 103, ErrMessage: "required param is empty"}
	ParamUnknownActionTypeError   = &DouyinError{ErrCode: 104, ErrMessage: "unknown action type"}
)

var (
	// AuthError 2xx
	AuthTokenFailError              = &DouyinError{ErrCode: 201, ErrMessage: "token is wrong or timeout"}
	AuthUsernameOrPasswordFailError = &DouyinError{ErrCode: 202, ErrMessage: "username or password is wrong"}
)

var (
	// DBError 3xx
	DBCreateUserError          = &DouyinError{ErrCode: 301, ErrMessage: "create user fail"}
	DBCreateVideoError         = &DouyinError{ErrCode: 302, ErrMessage: "create video fail"}
	DBCreateFavoriteEventError = &DouyinError{ErrCode: 303, ErrMessage: "create favorite event fail"}
	DBCreateCommentEventError  = &DouyinError{ErrCode: 304, ErrMessage: "create comment event fail"}
	DBCreateRelationEventError = &DouyinError{ErrCode: 305, ErrMessage: "create relation event fail"}
	DBDeleteFavoriteEventError = &DouyinError{ErrCode: 313, ErrMessage: "delete favorite event fail"}
	DBDeleteCommentEventError  = &DouyinError{ErrCode: 314, ErrMessage: "delete comment event fail"}
	DBDeleteRelationEventError = &DouyinError{ErrCode: 315, ErrMessage: "delete relation event fail"}
)

var (
	// uploadError 4xx
	UploadFileExistError = &DouyinError{ErrCode: 401, ErrMessage: "upload file exist"}
)

var (
	UnknownError = &DouyinError{ErrCode: 999}
)

func HandleBindError(err error) *DouyinError {
	switch err.(type) {
	case validator.ValidationErrors:
		// only consider the first err message
		errMessage := err.(validator.ValidationErrors)[0].Error()
		switch {
		case strings.HasSuffix(errMessage, "'required' tag"):
			return ParamEmptyError
		case strings.HasSuffix(errMessage, "'oneof' tag"):
			return ParamUnknownActionTypeError
		case strings.HasSuffix(errMessage, "'lte' tag"):
			return ParamInputLengthExceededError
		default:
			dyerr := UnknownError
			dyerr.ErrMessage = errMessage
			return dyerr
		}
	case *strconv.NumError:
		return ParamInputTypeError
	default:
		fmt.Printf("%T\n", err)
		dyerr := UnknownError
		dyerr.ErrMessage = err.Error()
		return dyerr
	}
}
