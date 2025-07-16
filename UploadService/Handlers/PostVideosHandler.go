package handlers

import (
	"errors"
	"math"
	"uploadservice/Errors/Handlers/PostVideosHandler"
	"uploadservice/Errors/Services/FileSystemService"
	"uploadservice/Interfaces/Services"
	"uploadservice/Interfaces/Services/FileSystemService"
	"uploadservice/Models/PostVideosModels"
)

type PostVideosHandler struct {
	FileSystemService filesystemInterfaces.IFileSystemService
	HashingService    interfaces.IHashingService
}

func (handler PostVideosHandler) Handle(request postVideosModels.PostVideosRequest) error {

	hashedName := handler.HashingService.Hash(request.Headers.FileName)

	_, err := handler.FileSystemService.GetFile(hashedName)

	if err == nil { // crazy I know
		return postVideosHandlerErrors.FileExistsError{FileName: request.Headers.FileName}
	}

	if err != nil {
		if !errors.As(err, &filesystemServiceErrors.FileNotFoundError{}) {
			return err
		}

		return err
	}

	content := make([]byte, int(math.Min(float64(request.Headers.ContentLength), math.Pow(10, 6))))

	if request.Headers.ContentLength > 0 {

		defer request.Content.Close()

		request.Content.Read(content)

		// Do something with content and error here

	}

	err = handler.FileSystemService.CreateNewFile(hashedName, content)

	if err != nil {
		return postVideosHandlerErrors.FailedToCreateResource{FileName: request.Headers.FileName, ErrorMessage: err.Error()}
	}

	return nil

}
