package web

import (
	"encoding/json"
	"net/http"
)

// HeaderWrite write error to Header
func HeaderWrite(w *http.ResponseWriter, status int, err error) {
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(err.Error())
}
