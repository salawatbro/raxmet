package errors

import "errors"

var (
	ErrInvalidId = errors.New("invalid ID")
	ErrNotFound  = errors.New("not found")
	ErrInternal  = errors.New("internal error")
)
