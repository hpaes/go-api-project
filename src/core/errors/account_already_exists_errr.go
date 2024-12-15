package errors

import "fmt"

type AccountAlreadyExistsErr struct {
	email string
}

func NewAccountAlreadyExistsErr(email string) error {
	return AccountAlreadyExistsErr{email: email}
}

func (e AccountAlreadyExistsErr) Error() string {
	return fmt.Sprintf("Account with email %s already exists", e.email)
}
