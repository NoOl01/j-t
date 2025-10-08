package errs

import "errors"

var (
	UserNotFound         = errors.New("user not found")
	UserAlreadyExist     = errors.New("user already exist")
	WrongPassword        = errors.New("wrong password")
	WrongToken           = errors.New("wrong token")
	InvalidBody          = errors.New("invalid body")
	MissingAuthToken     = errors.New("authorization token is missing")
	InvalidAuthToken     = errors.New("invalid authorization token")
	WrongAuthTokenFormat = errors.New("wrong authorization token format")
	ProductZeroCount     = errors.New("item count must be greater than zero")
)
