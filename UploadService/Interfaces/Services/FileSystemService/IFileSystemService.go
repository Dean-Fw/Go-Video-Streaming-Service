package filesystemInterfaces

import "io"

type IFileSystemService interface {
	CreateNewFile(fileNameHash string, content []byte) error
	GetFile(fileNameHash string) (file io.WriteCloser, err error)
	UpdateFile(file io.WriteCloser, content []byte) (err error)
}
