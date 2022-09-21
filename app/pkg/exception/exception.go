package exception

import "net/http"

type exception struct {
	msg      string
	httpCode int
}

var BadRequest error
var InternalError error
var UnauthorizedError error
var ForbiddenError error

func init() {
	BadRequest = exception{"BadRequest", http.StatusBadRequest}

	InternalError = exception{"InternalServerError", http.StatusInternalServerError}

	UnauthorizedError = exception{"Unauthorized", http.StatusUnauthorized}

	ForbiddenError = exception{"Forbidden", http.StatusOK}
}

func (e exception) Error() string {
	return e.msg
}

func _(httpCode int, msg string) error {
	return exception{
		msg:      msg,
		httpCode: httpCode,
	}
}

func GetHttpStatusCode(e error) int {

	switch e.Error() {
	case InternalError.Error():
		return InternalError.(exception).httpCode
	case UnauthorizedError.Error():
		return UnauthorizedError.(exception).httpCode
	case ForbiddenError.Error():
		return ForbiddenError.(exception).httpCode
	case BadRequest.Error():
		return BadRequest.(exception).httpCode
	}

	return http.StatusOK
}
