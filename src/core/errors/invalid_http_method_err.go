package errors

import "fmt"

type InvalidHttpMethodErr struct {
	method string
}

func NewInvalidHttpMethodErr(method string) error {
	return InvalidHttpMethodErr{method: method}
}

func (e InvalidHttpMethodErr) Error() string {
	return fmt.Sprintf("Invalid HTTP method %s", e.method)
}
