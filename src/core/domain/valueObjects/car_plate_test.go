package valueObjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCarPlate(t *testing.T) {
	carPlate, err := NewCarPlate("ABC-1B34")
	assert.NoError(t, err)
	assert.Equal(t, "ABC-1B34", carPlate.Value)
}

func TestInvalidCarPlates(t *testing.T) {
	tests := []struct {
		carPlate string
	}{
		{"ABC-1B3"},
		{"ABC-1B345"},
		{"ABC-113A"},
		{"ABC-1111"},
		{"234-1B34"},
	}
	for _, test := range tests {
		carPlate, err := NewCarPlate(test.carPlate)
		assert.Nil(t, carPlate)
		assert.Errorf(t, err, "Invalid car plate")
	}
}
