package valueObjects

import (
	"regexp"

	"github.com/pkg/errors"
)

type Name struct {
	Value string
}

func NewName(value string) (*Name, error) {
	if !isValidName(value) {
		return nil, errors.New("Invalid name")
	}
	return &Name{Value: value}, nil
}

func isValidName(value string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z]+ [a-zA-Z]+$`)
	return regex.MatchString(value)
}
