package valueObjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCpf(t *testing.T) {
	cpf, err := NewCpf("123.456.789-09")
	assert.NoError(t, err)
	assert.NotNil(t, cpf)
	assert.Equal(t, "123.456.789-09", cpf.Value)
}

func TestInvalidCpfs(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"123.456.789-10"},
		{"111.111.111-11"},
		{"123.456.789-"},
		{"123.456.789-0"},
		{"123.456.789-!"},
		{""},
	}
	for _, test := range tests {
		cpf, err := NewCpf(test.input)
		assert.Nil(t, cpf)
		assert.Errorf(t, err, "Invalid cpf")
	}
}
