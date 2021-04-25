package errs

import "errors"

// Module-level error
var (
	ErrInvalidArgument        = errors.New("invalid argument")
	ErrResponseInvalidData    = errors.New("invalid data")
	ErrResponseInvalidContent = errors.New("invalid content")
	ErrRequestFailed          = errors.New("request failed")
	ErrUnknown                = errors.New("unknown error")
)