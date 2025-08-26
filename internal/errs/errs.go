package errs

import "errors"

var (
	UserNotFound     = errors.New("user not found")
	UserAlreadyExist = errors.New("user already exist")
	WrongPassword    = errors.New("wrong password")
	WrongToken       = errors.New("wrong token")
	InvalidBody      = errors.New("invalid body")
	HelperEmptyValue = errors.New("empty value")
)
