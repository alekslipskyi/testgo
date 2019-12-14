package Router

import (
	"core/logger"
	"encoding/json"
	"helpers/errors"
	"net/http"
	"reflect"
)

type failedHandler func(error errors.IRequestError)
type sendJsonHandler func(model interface{}, statusCode int)
type sendHandler func(message string, statusCode int)

var logConn = logger.Logger{"Router Connection"}

type connection struct {
	closed bool
	w      http.ResponseWriter
	r      *http.Request
}

func (conn *connection) sendJson(model interface{}, statusCode int) {
	if conn.closed {
		logConn.Warn(twiceCallError)
		return
	}
	conn.closed = true
	conn.w.WriteHeader(statusCode)
	conn.w.Header().Set("Content-Type", "application/json")

	payload, _ := json.Marshal(&model)

	if string(payload) == "null" {
		switch reflect.TypeOf(model).Kind() {
		case reflect.Slice:
			payload = []byte("[]")
		}
	}

	_, _ = conn.w.Write(payload)
}

func (conn *connection) send(message string, statusCode int) {
	if conn.closed {
		logConn.Warn(twiceCallError)
		return
	}

	conn.closed = true

	conn.w.WriteHeader(statusCode)
	_, _ = conn.w.Write([]byte(message))
}

func (conn *connection) reject(error errors.IRequestError) {
	if conn.closed {
		logConn.Warn(twiceCallError)
		return
	}
	conn.closed = true

	conn.w.Header().Set("Content-Type", "application/json")
	conn.w.WriteHeader(error.StatusCode)

	js, _ := json.Marshal(error)
	_, _ = conn.w.Write(js)
}
