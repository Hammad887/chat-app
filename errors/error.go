package errors

import "errors"

var (
	// ErrNoContent custom error for no record found
	ErrNoContent = errors.New("no model exists in database")
	// ErrConflict custom error for record already exist
	ErrConflict = errors.New("model already exists in database")
)
