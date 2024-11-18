package constants

import "errors"

var (
	ErrInvalidId              = errors.New("invalid ID")
	ErrNotFound               = errors.New("not found")
	ErrInternal               = errors.New("internal error")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrTokenExpired           = errors.New("token expired")
	ErrSomethingWentWrong     = errors.New("something went wrong")
	ErrEmailExists            = errors.New("email exists")
)
