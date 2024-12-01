package valueObjects

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

type Cpf struct {
	Value string
}

func NewCpf(value string) (*Cpf, error) {
	if !validate(value) {
		return nil, errors.New("Invalid cpf")
	}
	return &Cpf{Value: value}, nil
}
func validate(cpf string) bool {
	if cpf == "" {
		return false
	}
	cpf = sanitize(cpf)
	if isInvalidLength(cpf) {
		return false
	}
	if allDigitsAreEqual(cpf) {
		return false
	}
	dg1 := calculateDigit(cpf, 10)
	dg2 := calculateDigit(cpf, 11)
	return extractCheckDigits(cpf) == strconv.Itoa(dg1)+strconv.Itoa(dg2)
}

func extractCheckDigits(cpf string) string {
	return cpf[9:]
}

func calculateDigit(cpf string, weight int) int {
	sum := 0
	for _, digit := range cpf {
		if weight > 1 {

			sum += int(digit-'0') * (weight)
			weight--
		}
	}
	rest := sum % 11
	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func allDigitsAreEqual(cpf string) bool {
	firstChar := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != firstChar {
			return false
		}
	}
	return true
}

func isInvalidLength(cpf string) bool {
	return len(cpf) != 11
}

func sanitize(cpf string) string {
	re, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(cpf, "")
}
