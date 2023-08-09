package httperror

import (
	"log"
	"net/http"
)

// Error : Response with given status code
func Error(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

// BadRequest
// Response with Bad Request (400)
func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// InternalServerError
// Response with InternalServerError (500)
func InternalServerError(w http.ResponseWriter, err error) {
	log.Printf("[InternalServerError] %+v", err)

	Error(w, http.StatusInternalServerError)
}
