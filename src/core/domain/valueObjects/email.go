package valueObjects

import (
	"regexp"

	"github.com/pkg/errors"
)

type Email struct {
	Value string
}

func NewEmail(value string) (*Email, error) {
	if !isValidEmail(value) {
		return nil, errors.New("Invalid email")
	}
	return &Email{Value: value}, nil
}

func isValidEmail(value string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(value)
}
