package domain

import "errors"

var (
	// ErrInternalServerError Will throw if any the internal server error happen.
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists.
	ErrNotFound = errors.New("your requested item was not found")
	// ErrBadParamInput will throw if the given request-body or params is not valid.
	ErrBadParamInput = errors.New("given param is not valid")
)
