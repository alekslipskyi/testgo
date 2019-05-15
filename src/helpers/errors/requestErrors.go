package errors

type RequestError struct {
	StatusCode int
	Message    string
	Token      string
}
