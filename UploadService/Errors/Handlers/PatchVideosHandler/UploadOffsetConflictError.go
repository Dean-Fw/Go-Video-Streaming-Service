package patchVideosHandlerErrors

import "fmt"

type UploadOffsetConflictError struct {
	ResourceOffset int64
	RequestOffset  int64
}

func (err UploadOffsetConflictError) Error() string {
	return fmt.Sprintf(
		"Cannot upload new bytes to file as the provided offset of '%dB' is before the offset of the resource '%dB'", err.RequestOffset, err.ResourceOffset)
}
