package response

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")
)

type Error struct {
	Message string
	Code    string
}

func NewError(message string, code string) Error {
	return Error{
		Message: message,
		Code:    code,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "999999")
	ErrorBadRequest = NewError("bad request", "40000")
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001")
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002")
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003")
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004")
	ErrorAuthIsNotExists       = NewError(ErrAuthIsNotExists.Error(), "40401")
	ErrorEmailAlreadyUsed      = NewError(ErrEmailAlreadyUsed.Error(), "40901")
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40101")
)

var (
	ErrorMapping = map[string]Error{
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error():       ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
	}
)
