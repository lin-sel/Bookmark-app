package web

// HTTPError Represent Send Back as Response
type HTTPError struct {
	HTTPStatus int
	ErrorKey   string
}

func (err HTTPError) Error() string {
	return err.ErrorKey
}

// NewHTTPError return New Instance of HTTPError
func NewHTTPError(err string, statuscode int) HTTPError {
	return HTTPError{
		HTTPStatus: statuscode,
		ErrorKey:   err,
	}
}
