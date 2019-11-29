package channel

import (
	"helpers/errors"
	"net/http"
)

var CHANNEL_ALREADY_EXISTS = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "Channel already exists",
	Token:      "CHANNEL_ALREADY_EXISTS",
}

var USER_ALREADY_INVITED = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "User already in channel",
	Token:      "USER_ALREADY_INVITED",
}

var USER_ID_IS_VALID = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "User id or channel id is not valid",
	Token:      "USER_ID_IS_VALID",
}
