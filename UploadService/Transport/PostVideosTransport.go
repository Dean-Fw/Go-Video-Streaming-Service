package transport

import (
	"net/http"
	"uploadservice/Errors/Transport"
	"uploadservice/Interfaces/Handlers"
	"uploadservice/Interfaces/Services"
	"uploadservice/Models/PostVideosModels"
)

type PostVideosTransport struct {
	HeaderValidatorService interfaces.IHeaderValidatorService
	PostVideosHandler      handlerInterfaces.IPostVideosHandler
}

func (transport PostVideosTransport) Receive(w http.ResponseWriter, r *http.Request) {
	headers, err := transport.checkHeaders(r.Header)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	transport.PostVideosHandler.Handle(postVideosModels.PostVideosRequest{Headers: headers, Content: r.Body})

}

func (transport PostVideosTransport) checkHeaders(headers http.Header) (postVideosModels.PostUploadStartHeaders, error) {
	uploadLengthHeader, err := transport.HeaderValidatorService.CheckHeaderInt64Value("Upload-Length", headers.Get("Upload-Length"))

	if err != nil {
		return postVideosModels.PostUploadStartHeaders{}, transportErrors.BadRequestError{Message: err.Error()}
	}

	contentLengthHeader, err := transport.HeaderValidatorService.CheckHeaderInt64Value("Content-Length", headers.Get("Content-Length"))

	if err != nil {
		return postVideosModels.PostUploadStartHeaders{}, transportErrors.BadRequestError{Message: err.Error()}
	}

	fileName, err := transport.HeaderValidatorService.CheckHeaderStringValue("File-Name", headers.Get("File-Name"))

	if err != nil {
		return postVideosModels.PostUploadStartHeaders{}, transportErrors.BadRequestError{Message: err.Error()}
	}

	return postVideosModels.PostUploadStartHeaders{
		UploadLength:  uploadLengthHeader,
		FileName:      fileName,
		ContentLength: contentLengthHeader,
	}, nil
}
