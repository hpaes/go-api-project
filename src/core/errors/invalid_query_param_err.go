package errors

import "fmt"

type InvalidQueryParamPayloadErr struct {
	query string
}

func NewInvalidQueryParamErr(query string) error {
	return InvalidQueryParamPayloadErr{query: query}
}

func (e InvalidQueryParamPayloadErr) Error() string {
	return fmt.Sprintf("Invalid query param: %s is missing", e.query)
}
