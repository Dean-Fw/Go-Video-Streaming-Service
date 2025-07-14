package filesystemModels

import "io"

type GetFilesResponseModel struct {
	FileWriter io.WriteCloser
	FileSize   int64
}
