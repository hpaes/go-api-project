package errors

type InternalServerErr struct {
	message string
	err     error
}

func NewInternalServerErr(message string, err error) *InternalServerErr {
	return &InternalServerErr{
		message: message,
		err:     err,
	}
}

func (e *InternalServerErr) Error() string {
	return e.message
}
