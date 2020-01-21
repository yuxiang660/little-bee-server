// A wrapper for standard error package. Don't use standard error package directly in the project.

package errors

import (
	"net/http"
)

type response map[string]interface{}

// Error defines the interface to manager the error
type Error interface {
	Error() string
	ResponseBody() response
	StatusCode() int 
}

// new returns an internal error that has implemented Error interface.
func newInternalError(text string) Error {
	return &errorImpl{
		statusCode: http.StatusInternalServerError,
		message: text,
	}
}

// new returns an HTTP error that has implemented Error interface.
func newHTTPError(code int, text string) Error {
	return &errorImpl{
		statusCode: code,
		message: text,
	}
}

// errorImpl is a implementation of Error.
type errorImpl struct {
	statusCode int
	message string
}

func (e *errorImpl) Error() string {
	return e.message
}

func (e *errorImpl) ResponseBody() response {
	return response{
		"code": e.statusCode,
		"message": e.message,
	}
}

func (e *errorImpl) StatusCode() int {
	return e.statusCode
}
