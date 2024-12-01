package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidQueryParamErr(t *testing.T) {
	var query = "test query"
	test := struct {
		name  string
		query string
		err   string
	}{
		name:  "Test InvalidRequestPayloadErr",
		query: query,
		err:   fmt.Sprintf("Invalid query param: %s is missing", query),
	}

	t.Run(test.name, func(t *testing.T) {
		err := NewInvalidQueryParamErr(test.query)
		assert.EqualError(t, err, test.err)
	})
}
