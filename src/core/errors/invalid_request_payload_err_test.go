package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequestPayloadErr(t *testing.T) {
	var message = "test message"
	test := struct {
		name    string
		message string
		err     string
	}{
		name:    "Test InvalidRequestPayloadErr",
		message: message,
		err:     fmt.Sprintf("Invalid request payload: %s", message),
	}

	t.Run(test.name, func(t *testing.T) {
		err := NewInvalidRequestPayloadErr(test.message)
		assert.EqualError(t, err, test.err)
	})
}
