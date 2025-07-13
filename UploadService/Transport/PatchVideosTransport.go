package transport

import (
	"fmt"
	"net/http"
	"uploadservice/Errors"
	"uploadservice/Interfaces/Handlers"
	"uploadservice/Interfaces/Services"
	"uploadservice/Models/PatchVideosModels"
)

type PatchVideosTransport struct {
	HeaderValidatorService interfaces.IHeaderValidatorService
	PatchVideosHandler     handlerInterfaces.IPatchVideosHandler
}

func (transport PatchVideosTransport) Receive(w http.ResponseWriter, r *http.Request) {
	headers, err := transport.checkHeaders(r.Header)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := transport.PatchVideosHandler.Handle(models.PatchVideosRequest{Resource: r.PathValue("id"), Headers: headers, Content: r.Body})

	fmt.Fprintf(w, "%s", resp)
}

func (transport PatchVideosTransport) checkHeaders(headers http.Header) (models.PatchVideosHeaders, error) {
	contentType, err := transport.HeaderValidatorService.CheckHeaderStringValue("Content-Type", headers.Get("Content-Type"))

	if err != nil {
		return models.PatchVideosHeaders{}, errors.BadRequestError{Message: err.Error()}
	}

	contentLength, err := transport.HeaderValidatorService.CheckHeaderIntValue("Content-Length", headers.Get("Content-Length"))

	if err != nil {
		return models.PatchVideosHeaders{}, errors.BadRequestError{Message: err.Error()}
	}

	uploadOffset, err := transport.HeaderValidatorService.CheckHeaderIntValue("Upload-Offset", headers.Get("Upload-Offset"))

	if err != nil {
		return models.PatchVideosHeaders{}, errors.BadRequestError{Message: err.Error()}
	}

	return models.PatchVideosHeaders{
		UploadOffset:  uploadOffset,
		ContentLength: contentLength,
		ContentType:   contentType,
	}, nil
}
