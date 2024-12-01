package errors

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountNotFoundErr(t *testing.T) {
	var accountId = uuid.New().String()
	test := struct {
		name      string
		accountId string
		err       string
	}{
		name:      "Test AccountNotFoundErr",
		accountId: accountId,
		err:       fmt.Sprintf("Account with id %s not found", accountId),
	}

	t.Run(test.name, func(t *testing.T) {
		err := NewAccountNotFoundErr(test.accountId)
		assert.EqualError(t, err, test.err)
	})
}
