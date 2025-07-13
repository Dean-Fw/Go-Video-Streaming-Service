package models

import "io"

type PatchVideosRequest struct {
	Resource string
	Headers  PatchVideosHeaders
	Content  io.ReadCloser
}
