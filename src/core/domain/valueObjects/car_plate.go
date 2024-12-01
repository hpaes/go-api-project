package valueObjects

import (
	"regexp"

	"github.com/pkg/errors"
)

type CarPlate struct {
	Value string
}

func NewCarPlate(value string) (*CarPlate, error) {
	if !isValidCarPlate(value) {
		return nil, errors.New("Invalid car plate")
	}

	return &CarPlate{Value: value}, nil
}

func isValidCarPlate(carPlate string) bool {
	regex, _ := regexp.Compile(`^[A-Z]{3}-[0-9][A-Z][0-9]{2}$`)
	return regex.MatchString(carPlate)
}
