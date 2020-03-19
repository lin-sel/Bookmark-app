package web

import (
	"encoding/json"
	"net/http"
)

// WriteErrorResponse write error to Header
func WriteErrorResponse(w *http.ResponseWriter, httperror HTTPError) {
	(*w).WriteHeader(httperror.HTTPStatus)
	json.NewEncoder(*w).Encode(httperror.Error())
}
