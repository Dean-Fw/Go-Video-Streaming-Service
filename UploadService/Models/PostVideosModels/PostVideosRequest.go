package postVideosModels

import "io"

type PostVideosRequest struct {
	Headers PostUploadStartHeaders
	Content io.ReadCloser
}
