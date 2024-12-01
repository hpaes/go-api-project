package errors

import "fmt"

type InvalidRequestPayloadErr struct {
	message string
}

func NewInvalidRequestPayloadErr(message string) error {
	return InvalidRequestPayloadErr{message: message}
}

func (e InvalidRequestPayloadErr) Error() string {
	return fmt.Sprintf("Invalid request payload: %s", e.message)
}
