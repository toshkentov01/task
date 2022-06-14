package errs

import "errors"

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("not found")

	// ErrInternal ...
	ErrInternal = errors.New("internal server error")
)