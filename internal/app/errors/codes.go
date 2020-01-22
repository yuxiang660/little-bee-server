package errors

import (
	"net/http"
)

// Error with HTTP status code.
// If the error is an internal error, the status code is http.StatusInternalServerError.
// If the error is an internal error, the status code refers to "net/http/status.go".
var (
	ErrInvalidToken = newInternalError("Invalid Token")
	ErrUnknowDatabase = newInternalError("Unknow Database")

	NoError = newHTTPError(http.StatusOK, "OK")
	ErrNotFound = newHTTPError(http.StatusNotFound, "Not Found")
	ErrMethodNotAllowed = newHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
	ErrUnauthorized = newHTTPError(http.StatusUnauthorized, "Non-Authoritative Information")
	ErrInternalServerError = newHTTPError(http.StatusInternalServerError, "Internal Server Error")
)
