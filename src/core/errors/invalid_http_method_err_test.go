package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidHttpMethodErr(t *testing.T) {
	var method = "GET"
	test := struct {
		name   string
		method string
		err    string
	}{
		name:   "Test AccountNotFoundErr",
		method: method,
		err:    fmt.Sprintf("Invalid HTTP method %s", method),
	}

	t.Run(test.name, func(t *testing.T) {
		err := NewInvalidHttpMethodErr(test.method)
		assert.EqualError(t, err, test.err)
	})
}
