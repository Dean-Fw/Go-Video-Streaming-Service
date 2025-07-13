package errors

import "fmt"

type HeaderIsNotIntError struct {
	HeaderName string
}

func (err HeaderIsNotIntError) Error() string {
	return fmt.Sprintf("Header '%s' is not parsable as an interger", err.HeaderName)
}
