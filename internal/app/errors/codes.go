package errors

import (
	"net/http"
)

// Error with HTTP status code.
// If the error is an internal error, the status code is http.StatusInternalServerError.
// If the error is an internal error, the status code refers to "net/http/status.go".
var (
	ErrInvalidToken = newInternalError("Invalid Token")
	ErrUnknownDatabase = newInternalError("Unknown Database")

	NoError = newHTTPError(http.StatusOK, "OK")
	ErrNotFound = newHTTPError(http.StatusNotFound, "Not Found")
	ErrMethodNotAllowed = newHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
	ErrUnauthorized = newHTTPError(http.StatusUnauthorized, "Non-Authoritative Information")
	ErrInternalServerError = newHTTPError(http.StatusInternalServerError, "Internal Server Error")
	ErrBadRequestParam = newHTTPError(http.StatusBadRequest, "Bad Request Parameters")
	ErrInvalidUsername = newHTTPError(http.StatusBadRequest, "Invalid Username")
	ErrInvalidPassword = newHTTPError(http.StatusBadRequest, "Invalid Password")
	ErrInvalidUser = newHTTPError(http.StatusBadRequest, "Invalid User")
)
