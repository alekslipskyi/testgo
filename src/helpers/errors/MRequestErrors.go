package errors

type IRequestError struct {
	StatusCode int
	Message    string
	Token      string
}
