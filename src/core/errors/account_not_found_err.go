package errors

import "fmt"

// AccountNotFoundErr is a struct that represents an account not found error
type AccountNotFoundErr struct {
	accountId string
}

func NewAccountNotFoundErr(accountId string) error {
	return AccountNotFoundErr{accountId: accountId}
}

func (e AccountNotFoundErr) Error() string {
	return fmt.Sprintf("Account with id %s not found", e.accountId)
}
