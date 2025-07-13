package patchVideosHandlerErrors

import "fmt"

type ResourceNotFoundError struct {
	Resource string
}

func (err ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource '%s' not found", err.Resource)
}
