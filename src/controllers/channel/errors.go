package channel

import (
	"helpers/errors"
	"net/http"
)

var CHANNEL_NOT_FOUND = errors.IRequestError{
	StatusCode: http.StatusNotFound,
	Message:    "Channel doesn't exist",
	Token:      "CHANNEL_NOT_FOUND",
}

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

var NOT_ALLOWED_TO_DROP = errors.IRequestError{
	StatusCode: http.StatusForbidden,
	Message:    "You haven't permissions to delete this channel",
	Token:      "BAD_PERMISSION_DROP",
}

var NOT_ALLOWED_TO_INVITE = errors.IRequestError{
	StatusCode: http.StatusForbidden,
	Message:    "You haven't permissions to invite this channel",
	Token:      "BAD_PERMISSIONS_INVITE",
}

var USER_ID_IS_VALID = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "User id or channel id is not valid",
	Token:      "USER_ID_IS_VALID",
}
