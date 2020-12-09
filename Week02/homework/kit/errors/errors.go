package errors

import (
	"errors"
	"fmt"
)

type errResourceNotFound struct {
	cause error
}

func (e *errResourceNotFound) Cause() error {
	return e.cause
}

func (e *errResourceNotFound) Error() string {
	return fmt.Sprintf("Resource not found: %v", e.cause)
}

func IsErrResourceNotFound(err error) bool {
	var target *errResourceNotFound
	return errors.As(err, &target)
}

func NewErrResourceNotFound(cause error) error {
	return &errResourceNotFound{cause}
}
