package handlers

import (
	"log"
	"net/http"
	"strconv"
	"uploadservice/Errors"
	"uploadservice/Interfaces"
	"uploadservice/Models"
)

type PostStartUploadHandler struct {
	FileSystemService interfaces.IFileSystemService
	HashingService    interfaces.IHashingService
}

func (handler PostStartUploadHandler) Handle(w http.ResponseWriter, r *http.Request) {
	requiredHeaders, err := GetHeaders(r.Header)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	content := make([]byte, requiredHeaders.ContentLength)

	_, err = r.Body.Read(content)

	if err != nil {
		http.Error(w, "Failed to read content of request", 500)
		return
	}

	hashedName := handler.HashingService.Hash(requiredHeaders.FileName)

	err = handler.FileSystemService.CreateNewFile(hashedName, content)

	if err != nil {
		http.Error(w, "Failed to write new file", 500)
		return
	}

	log.Print("INFO: WRITTEN FILE")

	AddResponseHeaders(w, hashedName, len(content))

}

func GetHeaders(headers http.Header) (models.PostUploadStartHeaders, error) { // TODO Refactor this to be DRY
	uploadLengthHeader := headers.Get("Upload-Length")

	if uploadLengthHeader == "" {
		return models.PostUploadStartHeaders{}, errors.BadRequestError{Message: "Missing Upload-Length Header"}
	}

	uploadLength, err := strconv.Atoi(uploadLengthHeader)

	if err != nil {
		return models.PostUploadStartHeaders{}, errors.BadRequestError{Message: "Could not parse Upload-Length as an int"}
	}

	contentLengthHeader := headers.Get("Content-Length")

	if contentLengthHeader == "" {
		return models.PostUploadStartHeaders{}, errors.BadRequestError{Message: "Missing Content-Length Header"}
	}

	contentLength, err := strconv.Atoi(contentLengthHeader)

	if err != nil {
		return models.PostUploadStartHeaders{}, errors.BadRequestError{Message: "Could not parse Content-Length as an int"}
	}

	fileName := headers.Get("File-Name")

	if fileName == "" {
		return models.PostUploadStartHeaders{}, errors.BadRequestError{Message: "Missing File-Name Header"}
	}

	return models.PostUploadStartHeaders{
		UploadLength:  uploadLength,
		FileName:      fileName,
		ContentLength: contentLength,
	}, nil
}

func AddResponseHeaders(w http.ResponseWriter, hashedFileName string, uploadOffset int) {
	w.Header().Add("Location", hashedFileName)
	w.Header().Add("Upload-Offset", strconv.Itoa(uploadOffset))
}
