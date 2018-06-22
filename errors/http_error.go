package errors

func NewHTTPError(status int, code int, message string) error {
	return &HTTPError{status, code, message}
}

type HTTPError struct {
	status  int
	code    int
	message string
}

func (e *HTTPError) Status() int {
	return e.status
}

func (e *HTTPError) Code() int {
	return e.code
}

func (e *HTTPError) Error() string {
	return e.message
}
