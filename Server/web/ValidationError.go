package web

import "fmt"

// ValidationError error represent error occur in Validation
type ValidationError struct {
	ErrorKey string            `json:"errorKey"`
	Errors   map[string]string `json:"errors"`
}

func (verror ValidationError) Error() string {
	return fmt.Sprintf("Error: [%s - %s]", verror.ErrorKey, verror.Errors)
}

// NewValidationError Return Instance of ValidationError.
func NewValidationError(err string, failedValidation map[string]string) *ValidationError {
	return &ValidationError{
		ErrorKey: err,
		Errors:   failedValidation,
	}
}
