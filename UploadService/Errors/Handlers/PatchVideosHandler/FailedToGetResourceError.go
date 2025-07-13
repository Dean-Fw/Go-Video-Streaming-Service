package patchVideosHandlerErrors

import "fmt"

type FailedToGetResourceError struct {
	Resource string
}

func (err FailedToGetResourceError) Error() string {
	return fmt.Sprintf("Failed to get resource '%s'", err.Resource)
}
