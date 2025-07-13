package filesystemServiceErrors

import "fmt"

type FileNotFoundError struct {
	File string
}

func (err FileNotFoundError) Error() string {
	return fmt.Sprintf("File '%s' not found", err.File)
}
