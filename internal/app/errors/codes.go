package errors

import (
	"net/http"
)

// Error with HTTP status code.
// If the error is an internal error, the status code is http.StatusInternalServerError.
// If the error is an internal error, the status code refers to "net/http/status.go".
var (
	ErrorInValidToken = newInternalError("Invalid Token")
	ErrorUnknowDatabase = newInternalError("Unknow Database")

	NoError = newHTTPError(http.StatusOK, "OK")
	ErrorNotFound = newHTTPError(http.StatusNotFound, "Not Found")
	ErrorMethodNotAllowed = newHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
	ErrorUnauthorized = newHTTPError(http.StatusUnauthorized, "Non-Authoritative Information")
)
