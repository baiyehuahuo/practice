package dyerror

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
	DBDeleteFavoriteEventError = &DouyinError{ErrCode: 313, ErrMessage: "delete favorite event fail"}
)

var (
	// uploadError 4xx
	UploadFileExistError = &DouyinError{ErrCode: 401, ErrMessage: "upload file exist"}
)

var (
	UnknownError = &DouyinError{ErrCode: 999}
)
