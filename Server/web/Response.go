package web

import (
	"encoding/json"
	"net/http"
)

// RespondErrorMessage Write Error To Respond Writer.
func RespondErrorMessage(w *http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, message)
}

// RespondJSON Set Respond with status code.
func RespondJSON(w *http.ResponseWriter, statuscode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statuscode, response)
}

func writeToHeader(w *http.ResponseWriter, statuscode int, payload interface{}) {
	(*w).WriteHeader(statuscode)
	(*w).Write(payload.([]byte))
}

// RespondError returns a validation error else
func RespondError(w *http.ResponseWriter, err error) {
	switch err.(type) {
	case ValidationError:
		RespondJSON(w, http.StatusBadRequest, err)
	case HTTPError:
		httpError := err.(HTTPError)
		RespondErrorMessage(w, httpError.HTTPStatus, httpError.ErrorKey)
	default:
		RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
	}
}
