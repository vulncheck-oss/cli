package util

import (
	"fmt"
)

type FlagError struct {
	// Note: not struct{error}: only *FlagError should satisfy error.
	err error
}

func (fe *FlagError) Error() string {
	return fe.err.Error()
}

func (fe *FlagError) Unwrap() error {
	return fe.err
}

func FlagErrorf(format string, args ...interface{}) error {
	return FlagErrorWrap(fmt.Errorf(format, args...))
}

// FlagErrorWrap FlagError returns a new FlagError that wraps the specified error.
func FlagErrorWrap(err error) error { return &FlagError{err} }
