package valueObjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidName(t *testing.T) {
	name, err := NewName("John Doe")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", name.Value)
}

func TestInvalidNames(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"i!nvalidname", false},
		{"@invalid.com", false},
		{"", false},
	}

	for _, test := range tests {
		name, err := NewName(test.input)
		assert.Nil(t, name)
		assert.Error(t, err)
	}
}
