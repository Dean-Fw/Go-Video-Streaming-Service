package interfaces

import "os"

type IFileSystemService interface {
	GetFile(fileName string) (*os.File, error)
}
