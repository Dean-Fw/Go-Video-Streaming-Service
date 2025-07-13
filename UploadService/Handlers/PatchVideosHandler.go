package handlers

import (
	"errors"
	"uploadservice/Errors/Handlers/PatchVideosHandler"
	"uploadservice/Errors/Services/FileSystemService"
	"uploadservice/Interfaces/Services/FileSystemService"
	"uploadservice/Models/PatchVideosModels"
)

type PatchVideosHandler struct {
	FileSystemService filesystemInterfaces.IFileSystemService
}

func (handler PatchVideosHandler) Handle(request models.PatchVideosRequest) error {
	file, err := handler.FileSystemService.GetFile(request.Resource)

	if err != nil {
		if errors.As(err, &filesystemServiceErrors.FileNotFoundError{}) {
			return patchVideosHandlerErrors.ResourceNotFoundError{Resource: request.Resource}
		} else {
			return patchVideosHandlerErrors.FailedToGetResourceError{Resource: request.Resource}
		}
	}

	defer file.Close()
	defer request.Content.Close()

	contentBytes := make([]byte, request.Headers.ContentLength)
	request.Content.Read(contentBytes)

	err = handler.FileSystemService.UpdateFile(file, contentBytes)

	return nil
}
