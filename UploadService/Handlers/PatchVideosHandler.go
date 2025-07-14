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
	resp, err := handler.FileSystemService.GetFile(request.Resource)

	if err != nil {
		if errors.As(err, &filesystemServiceErrors.FileNotFoundError{}) {
			return patchVideosHandlerErrors.ResourceNotFoundError{Resource: request.Resource}
		} else {
			return patchVideosHandlerErrors.FailedToGetResourceError{Resource: request.Resource}
		}
	}

	defer resp.FileWriter.Close()
	defer request.Content.Close()

	if request.Headers.UploadOffset < resp.FileSize || request.Headers.UploadOffset > resp.FileSize {
		return patchVideosHandlerErrors.UploadOffsetConflictError{ResourceOffset: resp.FileSize, RequestOffset: request.Headers.UploadOffset}
	}

	contentBytes := make([]byte, request.Headers.ContentLength)
	request.Content.Read(contentBytes)

	err = handler.FileSystemService.UpdateFile(resp.FileWriter, contentBytes)

	return nil
}
