package valueObjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidEmail(t *testing.T) {
	email, err := NewEmail("johnDoe@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, "johnDoe@gmail.com", email.Value)
}

func TestInvalidEmails(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"invalidemail"},
		{"invalid.com"},
		{""},
		{"invalid@com"},
	}

	for _, test := range tests {
		email, err := NewEmail(test.input)
		assert.Nil(t, email)
		assert.Error(t, err)
	}
}
