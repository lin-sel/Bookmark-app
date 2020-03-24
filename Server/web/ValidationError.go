package web

import "fmt"

// ValidationError error represent error occur in Validation
type ValidationError struct {
	ErrorKey string            `json:"errorKey"`
	Errors   map[string]string `json:"errors"`
}

func (verror ValidationError) Error() string {
	var statement string
	for key, value := range verror.Errors {
		statement = fmt.Sprintf(`{"%s":"%s"}`, key, value)
	}
	return statement
}

// NewValidationError Return Instance of ValidationError.
func NewValidationError(err string, failedValidation map[string]string) *ValidationError {
	return &ValidationError{
		ErrorKey: err,
		Errors:   failedValidation,
	}
}
