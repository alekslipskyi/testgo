package requestError

import (
	"helpers/errors"
	"net/http"
)

var UNAUTHORIZED = errors.IRequestError{
	StatusCode: http.StatusUnauthorized,
	Message:    "Unauthorized",
	Token:      "UNAUTHORIZED",
}

var UNEXPECTED_ERROR = errors.IRequestError{
	StatusCode: http.StatusInternalServerError,
	Message:    "Something went wrong",
	Token:      "SOMETHING_WENT_WRONG",
}

var WRONG_IP = errors.IRequestError{
	StatusCode: http.StatusUnauthorized,
	Message:    "Ip is not allowed",
	Token:      "IP_IS_NOT_ALLOWED",
}

var USER_ALREADY_EXIST = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "User already exist",
	Token:      "USER_ALREADY_EXISTS",
}

var INVALID_CREDENTIAL = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid credentials",
	Token:      "INVALID_CREDENTIALS",
}

var NOT_FOUND = errors.IRequestError{
	StatusCode: http.StatusNotFound,
	Message:    "Not found",
	Token:      "NOT_FOUND",
}

var SOMETHING_WENT_WRONG = errors.IRequestError{
	StatusCode: http.StatusInternalServerError,
	Message:    "Something went wrong",
	Token:      "SOMETHING_WENT_WRONG",
}
