package Router

import (
	"encoding/json"
	"fmt"
	"helpers/errors"
	"net/http"
)

type failedHandler func(error errors.RequestError)
type sendJsonHandler func(json []byte, statusCode int)
type sendHandler func(message string, statusCode int)

type connection struct {
	closed bool
	w      http.ResponseWriter
	r      *http.Request
}

func (conn *connection) sendJson(json []byte, statusCode int) {
	if conn.closed {
		fmt.Println("Warning: ", twiceCallError)
		return
	}
	conn.closed = true
	conn.w.WriteHeader(statusCode)
	conn.w.Header().Set("Content-Type", "application/json")

	_, _ = conn.w.Write(json)
}

func (conn *connection) send(message string, statusCode int) {
	if conn.closed {
		fmt.Println("Warning: ", twiceCallError)
		return
	}

	conn.closed = true

	conn.w.WriteHeader(statusCode)
	_, _ = conn.w.Write([]byte(message))
}

func (conn *connection) reject(error errors.RequestError) {
	fmt.Println("failed")
	if conn.closed {
		fmt.Println("Warning: ", twiceCallError)
		return
	}
	conn.closed = true

	conn.w.Header().Set("Content-Type", "application/json")
	conn.w.WriteHeader(error.StatusCode)

	js, _ := json.Marshal(error)
	_, _ = conn.w.Write(js)
}
