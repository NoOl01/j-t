package errs

import "errors"

var (
	UserNotFound     = errors.New("user not found")
	UserAlreadyExist = errors.New("user already exist")
	WrongPassword    = errors.New("wrong password")
	InvalidBody      = errors.New("invalid body")
)
