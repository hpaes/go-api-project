package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerErr(t *testing.T) {
	var message = "test message"
	var err = errors.New("test error")
	test := struct {
		name    string
		message string
		err     error
	}{
		name:    "Test InternalServerErr",
		message: message,
		err:     err,
	}

	t.Run(test.name, func(t *testing.T) {
		err := NewInternalServerErr(test.message, test.err)
		assert.EqualError(t, err, test.message)
	})
}
