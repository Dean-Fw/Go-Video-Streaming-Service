package errors

import "fmt"

type HeaderMissingError struct {
	HeaderName string
}

func (err HeaderMissingError) Error() string {
	return fmt.Sprintf("Header '%s' is missing from request", err.HeaderName)
}
