package errors

import (
	"errors"
	"fmt"
)

type ApplicationError struct {
	Code    int
	Context string
	ErrMsg  string
}

func (ae *ApplicationError) ToError() error {
	errorMessage := fmt.Sprintf(
		"Error: %s, Context: %s, Code: %s",
		ae.ErrMsg,
		ae.Context,
		ae.Code,
	)
	return errors.New(errorMessage)
}
