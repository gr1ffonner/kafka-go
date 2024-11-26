package domainerrors

import "errors"

var (
	ErrBadRequest = errors.New("invalid request")
	ErrInternal   = errors.New("internal error")
	ErrNotFound   = errors.New("not found")
)
