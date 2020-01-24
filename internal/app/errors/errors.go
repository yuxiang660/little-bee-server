// A wrapper for standard error package. Don't use standard error package directly in the project.

package errors

import (
	"net/http"
)

type response map[string]interface{}

// Error defines the interface to manager the error
type Error interface {
	Error() string
	Body() response
	Code() int 
}

// new returns an internal error that has implemented Error interface.
func newInternalError(text string) Error {
	return &impl{
		statusCode: http.StatusInternalServerError,
		message: text,
	}
}

// new returns an HTTP error that has implemented Error interface.
func newHTTPError(code int, text string) Error {
	return &impl{
		statusCode: code,
		message: text,
	}
}

// impl is a implementation of Error.
type impl struct {
	statusCode int
	message string
}

func (e *impl) Error() string {
	return e.message
}

func (e *impl) Body() response {
	return response{
		"code": e.statusCode,
		"message": e.message,
	}
}

func (e *impl) Code() int {
	return e.statusCode
}
