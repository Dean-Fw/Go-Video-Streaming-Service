package postVideosHandlerErrors

import "fmt"

type FailedToCreateResource struct {
	FileName     string
	ErrorMessage string
}

func (err FailedToCreateResource) Error() string {
	return fmt.Sprintf("Failed to create resource for file '%s' because: %s ", err.FileName, err.ErrorMessage)
}
