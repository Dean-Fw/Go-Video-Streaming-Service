package filesystemInterfaces

import (
	"io"
	"uploadservice/Models/Services/FileSystemService"
)

type IFileSystemService interface {
	CreateNewFile(fileNameHash string, content []byte) error
	GetFile(fileNameHash string) (response filesystemModels.GetFilesResponseModel, err error)
	UpdateFile(file io.WriteCloser, content []byte) (err error)
}
