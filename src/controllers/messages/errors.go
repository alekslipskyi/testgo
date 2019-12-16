package messages

import (
	"helpers/errors"
	"net/http"
)

var WRITE_IS_NOT_ALLOWED = errors.IRequestError{
	StatusCode: http.StatusBadRequest,
	Message:    "You haven't permissions to write to this channel",
	Token:      "BAD_PERMISSIONS_WRITE",
}
